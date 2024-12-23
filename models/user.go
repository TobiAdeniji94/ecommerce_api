package models

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// user model to represent user.
// Role can be "user" or "admin".
type User struct {
    ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
    Email    string `gorm:"unique;not null" json:"email"`
    Password string `json:"-"`
    Role     string `json:"role" gorm:"default:user"`
}

// BeforeCreate hook to generate a UUID for the user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    if u.ID == uuid.Nil {
        u.ID = uuid.New()
    }
    return
}

// LoginInput to bind the JSON body when a user logs in.
type LoginInput struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}
