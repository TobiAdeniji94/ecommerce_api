package models

import (
    "gorm.io/gorm"
)

// Product holds information about items available in the store.
type Product struct {
    gorm.Model
    Name        string  `gorm:"not null" json:"name"`
    Description string  `json:"description"`
    Price       float64 `gorm:"not null" json:"price"`
    Stock       int     `gorm:"not null" json:"stock"`
}
