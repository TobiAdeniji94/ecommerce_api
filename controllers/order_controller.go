package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"

    "github.com/TobiAdeniji94/ecommerce_api/config"
    "github.com/TobiAdeniji94/ecommerce_api/models"
)

// PlaceOrder allows an authenticated user to create a new order
func PlaceOrder(c *gin.Context) {
    // 1. Get the user UUID from JWT or context
    userData, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    userUUID, ok := userData.(uuid.UUID)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user UUID in context"})
        return
    }

    // 2. Parse the incoming order request
    var orderRequest struct {
        Items []struct {
            ProductID string `json:"product_id"` // We'll parse these as string -> uuid
            Quantity  int    `json:"quantity"`
        } `json:"items"`
    }
    if err := c.ShouldBindJSON(&orderRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 3. Create a new Order
    newOrder := models.Order{
        UserID: userUUID,
        Status: "Pending",
    }
    if err := config.DB.Create(&newOrder).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    // 4. Create OrderItems for each item
    for _, item := range orderRequest.Items {
        prodUUID, err := uuid.Parse(item.ProductID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product UUID"})
            return
        }

        orderItem := models.OrderItem{
            OrderID:   newOrder.ID,
            ProductID: prodUUID,
            Quantity:  item.Quantity,
        }
        if err := config.DB.Create(&orderItem).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order item"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order_id": newOrder.ID})
}

// GetUserOrders lists all orders for the authenticated user
func GetUserOrders(c *gin.Context) {
    userData, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    userUUID, ok := userData.(uuid.UUID)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user UUID in context"})
        return
    }

    var orders []models.Order
    // Preload "Items.Product" to fetch all order items + product data
    if err := config.DB.Preload("Items.Product").Where("user_id = ?", userUUID).Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": orders})
}

// CancelOrder cancels the order if it's still pending
func CancelOrder(c *gin.Context) {
    userData, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    userUUID, ok := userData.(uuid.UUID)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user UUID in context"})
        return
    }

    orderIDStr := c.Param("id")
    orderUUID, err := uuid.Parse(orderIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order UUID"})
        return
    }

    var order models.Order
    if err := config.DB.Where("id = ? AND user_id = ?", orderUUID, userUUID).First(&order).Error; err != nil {
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
    orderUUID, err := uuid.Parse(orderIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order UUID"})
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
    if err := config.DB.First(&order, "id = ?", orderUUID).Error; err != nil {
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
