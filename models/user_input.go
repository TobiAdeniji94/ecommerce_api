package models

// LoginInput to bind the JSON body when a user logs in.
type LoginInput struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}