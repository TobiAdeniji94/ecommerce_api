package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"

    "github.com/TobiAdeniji94/ecommerce_api/config"
    "github.com/TobiAdeniji94/ecommerce_api/models"
    "github.com/TobiAdeniji94/ecommerce_api/utils"
)

// RegisterUser handles user signup
func RegisterUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the password before saving
    hashedPassword, err := HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }
    input.Password = hashedPassword

    // Default role is user if not provided
    if input.Role == "" {
        input.Role = "user"
    }

    // Create the user in DB
    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!", "user_id": input.ID})
}

// LoginUser handles user authentication
func LoginUser(c *gin.Context) {
    var input models.LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Fetch user by email
    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Check password
    if err := CheckPassword(input.Password, user.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID.String(), user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token":   token,
        "message": "Login successful",
        "user_id": user.ID,
    })
}

// HashPassword hashes the plain text password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPassword compares plain password with hashed
func CheckPassword(plain, hashed string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
