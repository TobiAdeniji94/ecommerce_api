package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// Product holds information about items available in the store.
type Product struct {
    ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
    Name        string  `gorm:"not null" json:"name"`
    Description string  `json:"description"`
    Price       float64 `gorm:"not null" json:"price"`
    Stock       int     `gorm:"not null" json:"stock"`
}

// BeforeCreate hook to generate a UUID for the user
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
    if p.ID == uuid.Nil {
        p.ID = uuid.New()
    }
    return
}
