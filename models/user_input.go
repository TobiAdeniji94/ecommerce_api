package models

// UserInput represents the payload for creating a user.
type UserInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
    Role     string `json:"role" binding:"omitempty"`
}

// LoginInput to bind the JSON body when a user logs in.
type LoginInput struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}