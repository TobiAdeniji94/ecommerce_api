package models

// OrderItemInput represents an individual item in the order payload.
type OrderItemInput struct {
    ProductID string `json:"product_id binding:"required"` // Product ID
    Quantity  int    `json:"quantity binding:"required,min=1"`   // Quantity
}

// PlaceOrderInput represents the payload for placing an order.
type PlaceOrderInput struct {
    Items []OrderItemInput `json:"items"` // List of order items
}

// UpdateOrderStatusInput represents the payload for updating the order status.
type UpdateOrderStatusInput struct {
    Status string `json:"status" json:"items" binding:"required,dive` // order status
}