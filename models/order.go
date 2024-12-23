package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// Order represents a user's order. Contains multiple products via OrderItems.
type Order struct {
    ID     uuid.UUID   `gorm:"type:char(36);primaryKey" json:"id"`
    UserID uuid.UUID   `json:"user_id"`
    User   User        `gorm:"foreignKey:UserID" json:"user"` // Eager-loading reference
    Items  []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
    Status string      `gorm:"default:Pending" json:"status"` // Possible values: Pending, Completed, Canceled
}

// BeforeCreate hook to generate a UUID for the user
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
    if o.ID == uuid.Nil {
        o.ID = uuid.New()
    }
    return
}
