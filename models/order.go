package models

import (
    "gorm.io/gorm"
)

// Order represents a user's order. Contains multiple products via OrderItems.
type Order struct {
    gorm.Model
    UserID uint        `json:"user_id"`
    User   User        `gorm:"foreignKey:UserID" json:"user"` // Eager-loading reference
    Items  []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
    Status string      `gorm:"default:Pending" json:"status"` // Possible values: Pending, Completed, Canceled
}
