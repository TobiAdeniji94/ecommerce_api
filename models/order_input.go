package models

// OrderItemInput represents an individual item in the order payload.
type OrderItemInput struct {
    ProductID string `json:"product_id" binding:"required"`
    Quantity  int    `json:"quantity" binding:"required,min=1"`
}

// PlaceOrderInput represents the payload for placing an order.
type PlaceOrderInput struct {
    Items []OrderItemInput `json:"items" binding:"required,dive"`
}

// UpdateOrderStatusInput represents the payload for updating the order status.
type UpdateOrderStatusInput struct {
    Status string `json:"status" binding:"required"`
}
