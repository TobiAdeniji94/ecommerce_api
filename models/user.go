package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// User model represents a user in the system.
type User struct {
    ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
    Email     string    `gorm:"unique;not null" json:"email"`
    Password  string    `gorm:"type:varchar(255);not null" json:"password"`
    Role      string    `json:"role" gorm:"default:user"`
    CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate hook to generate a UUID for the user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    if u.ID == uuid.Nil {
        u.ID = uuid.New()
    }
    return
}
