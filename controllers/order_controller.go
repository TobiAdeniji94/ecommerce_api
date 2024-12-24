package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/TobiAdeniji94/ecommerce_api/config"
	"github.com/TobiAdeniji94/ecommerce_api/models"
)

// PlaceOrder allows an authenticated user to create a new order
// PlaceOrder godoc
// @Summary Place a new order
// @Description Allows an authenticated user to place an order with one or more products
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body models.PlaceOrderInput true "Order payload"
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse "Order created successfully"
// @Failure 400 {object} models.ValidationErrorResponse "Invalid order payload"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Failed to create order"
// @Router /orders [post]
func PlaceOrder(c *gin.Context) {
	userData, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Unauthorized"})
		return
	}

	userUUID, ok := userData.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid user ID"})
		return
	}

	var orderRequest models.PlaceOrderInput
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.ValidationErrorResponse{
			Errors: []models.ValidationError{
				{Field: "payload", Message: err.Error()},
			},
		})
		return
	}

	newOrder := models.Order{
		UserID: userUUID,
		Status: "Pending",
	}
	if err := config.DB.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to create order"})
		return
	}

	for _, item := range orderRequest.Items {
		prodUUID, err := uuid.Parse(item.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid product ID"})
			return
		}

		orderItem := models.OrderItem{
			OrderID:   newOrder.ID,
			ProductID: prodUUID,
			Quantity:  item.Quantity,
		}
		if err := config.DB.Create(&orderItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to create order item"})
			return
		}
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Order created successfully",
		Data:    gin.H{"order_id": newOrder.ID},
	})
}

// GetUserOrders lists all orders for the authenticated user
// GetUserOrders godoc
// @Summary Get all orders for a user
// @Description Retrieve a list of all orders placed by the authenticated user
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse "List of orders"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Failed to fetch orders"
// @Router /orders [get]
func GetUserOrders(c *gin.Context) {
	userData, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Unauthorized"})
		return
	}

	userUUID, ok := userData.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid user ID"})
		return
	}

	var orders []models.Order
	if err := config.DB.Preload("User").Preload("Items.Product").Where("user_id = ?", userUUID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to fetch orders"})
		return
	}

	for i := range orders {
		orders[i].User.Password = ""
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Orders retrieved successfully",
		Data:    orders,
	})
}

// CancelOrder cancels the order if it's still pending
// CancelOrder godoc
// @Summary Cancel an order
// @Description Allows an authenticated user to cancel an order if it is in "Pending" status
// @Tags Orders
// @Param id path string true "Order ID"
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse "Order canceled successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid order ID or status"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Order not found"
// @Failure 500 {object} models.ErrorResponse "Failed to cancel order"
// @Router /orders/{id}/cancel [put]
func CancelOrder(c *gin.Context) {
	userData, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Unauthorized"})
		return
	}

	userUUID, ok := userData.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid user ID"})
		return
	}

	orderIDStr := c.Param("id")
	orderUUID, err := uuid.Parse(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid order ID"})
		return
	}

	var order models.Order
	if err := config.DB.Where("id = ? AND user_id = ?", orderUUID, userUUID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Order not found"})
		return
	}

	if order.Status != "Pending" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Order cannot be canceled. Current status: " + order.Status,
		})
		return
	}

	order.Status = "Canceled"
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to cancel order"})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Order canceled successfully",
	})
}

// UpdateOrderStatus allows an admin to update an order status
// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Allows an admin to update the status of an order
// @Tags Orders
// @Param id path string true "Order ID"
// @Param status body models.UpdateOrderStatusInput true "Update order status payload"
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse "Order status updated successfully"
// @Failure 400 {object} models.ValidationErrorResponse "Invalid order ID or payload"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Order not found"
// @Failure 500 {object} models.ErrorResponse "Failed to update order status"
// @Router /orders/{id}/status [put]
func UpdateOrderStatus(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderUUID, err := uuid.Parse(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid order ID"})
		return
	}

	var requestBody models.UpdateOrderStatusInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, models.ValidationErrorResponse{
			Errors: []models.ValidationError{
				{Field: "status", Message: err.Error()},
			},
		})
		return
	}

	var order models.Order
	if err := config.DB.Preload("User").Preload("Items.Product").First(&order, "id = ?", orderUUID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Order not found"})
		return
	}

	order.Status = requestBody.Status
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to update order status"})
		return
	}

	order.User.Password = ""

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Order status updated successfully",
		Data:    order,
	})
}
