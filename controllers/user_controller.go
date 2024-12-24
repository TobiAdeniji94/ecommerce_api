package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/TobiAdeniji94/ecommerce_api/config"
	"github.com/TobiAdeniji94/ecommerce_api/models"
	"github.com/TobiAdeniji94/ecommerce_api/utils"
)

// RegisterUser handles user signup
// RegisterUser godoc
// @Summary Register a new user
// @Description Create a new user in the system with email, password, and optional role
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.UserInput true "User registration payload"
// @Success 200 {object} models.SuccessResponse "User registered successfully"
// @Failure 400 {object} models.ValidationErrorResponse "Validation errors"
// @Failure 500 {object} models.ErrorResponse "Failed to create user"
// @Router /users/register [post]
func RegisterUser(c *gin.Context) {
	var input models.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ValidationErrorResponse{
			Errors: []models.ValidationError{
				{Field: "payload", Message: err.Error()},
			},
		})
		return
	}

	// Validate input
	var validationErrors []models.ValidationError
	if input.Email == "" {
		validationErrors = append(validationErrors, models.ValidationError{
			Field:   "email",
			Message: "Email is required",
		})
	}
	if input.Password == "" {
		validationErrors = append(validationErrors, models.ValidationError{
			Field:   "password",
			Message: "Password is required",
		})
	}

	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, models.ValidationErrorResponse{
			Errors: validationErrors,
		})
		return
	}

	// Hash the password before saving
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Could not hash password",
		})
		return
	}

	// Map UserInput to User model
	user := models.User{
		Email:    input.Email,
		Password: hashedPassword,
		Role:     input.Role,
	}

	// Default role is user if not provided
	if user.Role == "" {
		user.Role = "user"
	}

	// Create the user in DB
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User registered successfully",
		Data:    gin.H{"user_id": user.ID},
	})
}

// LoginUser handles user authentication
// LoginUser godoc
// @Summary Authenticate a user
// @Description Authenticate a user with email and password, returning a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param login body models.LoginInput true "User login payload"
// @Success 200 {object} models.SuccessResponse "Login successful"
// @Failure 400 {object} models.ValidationErrorResponse "Validation errors"
// @Failure 401 {object} models.ErrorResponse "Invalid email or password"
// @Failure 500 {object} models.ErrorResponse "Failed to generate token"
// @Router /users/login [post]
func LoginUser(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ValidationErrorResponse{
			Errors: []models.ValidationError{
				{Field: "payload", Message: err.Error()},
			},
		})
		return
	}

	// Fetch user by email
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Invalid email or password",
		})
		return
	}

	// Check password
	if err := CheckPassword(input.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Invalid email or password",
		})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.String(), user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Login successful",
		Data: gin.H{
			"token":   token,
			"user_id": user.ID,
		},
	})
}

// HashPassword hashes the plain text password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword compares plain password with hashed
func CheckPassword(plain, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		log.Printf("Password comparison failed: %v\n", err)
	}
	return err
}
