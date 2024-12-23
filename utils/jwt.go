package utils

import (
    "fmt"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

// jwtKey to be initialized in the init() function
var jwtKey []byte

// init called when this package is imported.
func init() {
    keyFromEnv := os.Getenv("JWT_SECRET")
    if keyFromEnv == "" {
        panic("JWT_SECRET environment variable is not set or is empty")
    }
    jwtKey = []byte(keyFromEnv)
    fmt.Println("JWT_SECRET loaded successfully")
}

// GenerateJWT creates a new JWT token with the user ID, role, and a 24-hour expiration.
func GenerateJWT(userID uint, role string) (string, error) {
    // Create the token with claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "role":    role,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
        "iat":     time.Now().Unix(),
    })

    // Sign the token with our secret key
    return token.SignedString(jwtKey)
}

// GetJWTKey returns the loaded JWT key so other packages can use it.
func GetJWTKey() []byte {
    return jwtKey
}
