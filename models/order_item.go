package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// OrderItem represents a single product within an Order. 
type OrderItem struct {
    ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
    OrderID   uuid.UUID `json:"order_id"`
    ProductID uuid.UUID `json:"product_id"`
    Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
    Quantity  int       `json:"quantity"`
}

// BeforeCreate hook to generate a UUID for the user
func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
    if oi.ID == uuid.Nil {
        oi.ID = uuid.New()
    }
    return
}
