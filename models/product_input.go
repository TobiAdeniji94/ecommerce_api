package models

// ProductInput represents the payload for creating or updating a product.
type ProductInput struct {
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description" binding:"omitempty"`
    Price       float64 `json:"price" binding:"required,gt=0"`
    Stock       int     `json:"stock" binding:"required,min=0"`
}
