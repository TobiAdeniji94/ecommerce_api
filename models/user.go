package models

import (
    "gorm.io/gorm"
)

// user model to represent user.
// Role can be "user" or "admin".
type User struct {
    gorm.Model
    Email    string `gorm:"unique;not null" json:"email"`
    Password string `json:"-"`
    Role     string `json:"role" gorm:"default:user"`
}

// LoginInput to bind the JSON body when a user logs in.
type LoginInput struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}