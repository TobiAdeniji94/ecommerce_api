package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    
    "github.com/TobiAdeniji94/ecommerce_api/config"
    "github.com/TobiAdeniji94/ecommerce_api/models"
)

// PlaceOrder allows an authenticated user to create a new order
func PlaceOrder(c *gin.Context) {
    // 1. Get the user ID from JWT or context
    userID, exists := c.Get("userID") // Depends on your AuthMiddleware
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // 2. Parse the incoming order request
    var orderRequest struct {
        Items []struct {
            ProductID uint `json:"product_id"`
            Quantity  int  `json:"quantity"`
        } `json:"items"`
    }
    if err := c.ShouldBindJSON(&orderRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 3. Create a new Order
    newOrder := models.Order{
        UserID: userID.(uint),
        Status: "Pending",
    }
    if err := config.DB.Create(&newOrder).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    // 4. Create OrderItems for each item
    for _, item := range orderRequest.Items {
        orderItem := models.OrderItem{
            OrderID:   newOrder.ID,
            ProductID: item.ProductID,
            Quantity:  item.Quantity,
        }
        if err := config.DB.Create(&orderItem).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order item"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order created", "order_id": newOrder.ID})
}

// GetUserOrders lists all orders for the authenticated user
func GetUserOrders(c *gin.Context) {
    userID, exists := c.Get("userID") // depends on AuthMiddleware
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var orders []models.Order
    if err := config.DB.Preload("Items.Product").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": orders})
}

// CancelOrder cancels the order if it's still pending
func CancelOrder(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    orderIDStr := c.Param("id")
    orderID, err := strconv.Atoi(orderIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var order models.Order
    if err := config.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    if order.Status != "Pending" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Order cannot be canceled. Current status: " + order.Status})
        return
    }

    order.Status = "Canceled"
    if err := config.DB.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}

// UpdateOrderStatus allows an admin to update an order status
func UpdateOrderStatus(c *gin.Context) {
    // This route is admin-only (checked by AdminMiddleware)
    orderIDStr := c.Param("id")
    orderID, err := strconv.Atoi(orderIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var requestBody struct {
        Status string `json:"status"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var order models.Order
    if err := config.DB.First(&order, orderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    order.Status = requestBody.Status
    if err := config.DB.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully", "order": order})
}
