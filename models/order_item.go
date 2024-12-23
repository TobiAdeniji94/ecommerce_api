package models

import (
    "gorm.io/gorm"
)

// OrderItem represents a single product within an Order. 
type OrderItem struct {
    gorm.Model
    OrderID   uint    `json:"order_id"`
    ProductID uint    `json:"product_id"`
    Product   Product `gorm:"foreignKey:ProductID" json:"product"`
    Quantity  int     `json:"quantity"`
}
