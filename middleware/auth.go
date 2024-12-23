package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"

    "github.com/TobiAdeniji94/ecommerce_api/utils"
)

// AuthMiddleware checks for a valid JWT token in the Authorization header.
func AuthMiddleware(c *gin.Context) {
    // Example "Authorization" header: "Bearer <token>"
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
        c.Abort()
        return
    }

    // Strip out the "Bearer " part to get the token
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
        c.Abort()
        return
    }

    // Parse and validate the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Use the same signing key that was used to generate tokens
        return []byte(utils.GetJWTKey()), nil
    })

    // Error or invalid token
    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
    }

    // Extract custom claims from the token
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        // user_id was stored as float64, so we need to cast it to uint or int
        if userID, ok := claims["user_id"].(float64); ok {
            c.Set("userID", uint(userID))
        }
        if role, ok := claims["role"].(string); ok {
            c.Set("role", role)
        }
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        c.Abort()
        return
    }

    // Token is valid, proceed
    c.Next()
}

// AdminMiddleware ensures the user has the "admin" role.
func AdminMiddleware(c *gin.Context) {
    role, exists := c.Get("role")
    if !exists {
        c.JSON(http.StatusForbidden, gin.H{"error": "No role found"})
        c.Abort()
        return
    }

    // The role is stored as an interface{}; cast it to string
    roleStr, ok := role.(string)
    if !ok || roleStr != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient privileges"})
        c.Abort()
        return
    }

    // If role is admin, proceed
    c.Next()
}
