package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"

    "github.com/TobiAdeniji94/ecommerce_api/config"
    "github.com/TobiAdeniji94/ecommerce_api/routes"
)

func main() {
    // Load .env if it exists
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found or it failed to load. Continuing with system environment variables.")
    }

    // Connect to the database 
    config.ConnectDatabase()

    // new Gin router
    r := gin.Default()

    // Initialize routes
    routes.InitializeRoutes(r)

    // Fallback port if PORT is not set
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }

    // Run the server
    log.Printf("Server is running on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
