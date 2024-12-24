package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "github.com/TobiAdeniji94/ecommerce_api/docs"

    "github.com/TobiAdeniji94/ecommerce_api/config"
    "github.com/TobiAdeniji94/ecommerce_api/routes"
    "github.com/TobiAdeniji94/ecommerce_api/utils"
)

// @title E-Commerce API
// @version 1.0
// @description E-commerce API for managing orders and products.
// @host ecommerce-api-vkui.onrender.com
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Use "Bearer {your token}" to authorize
func main() {
    // Load .env if it exists
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found or it failed to load. Continuing with system environment variables.")
    }

    // Connect to database
    config.ConnectDatabase()

    // Gin router
    r := gin.Default()

    // CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "ecommerce-api-vkui.onrender.com"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                   
        AllowHeaders:     []string{"Authorization", "Content-Type"},                 
        ExposeHeaders:    []string{"Content-Length"},                                
        AllowCredentials: true,                                                     
        MaxAge:           12 * time.Hour,                                           
    }))

    // Rate Limiting middleware
    r.Use(utils.PerClientRateLimiter()) // Call the middleware from utils

    // Swagger docs
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Initialize routes
    routes.InitializeRoutes(r)

    // Fallback port if PORT is not set
    port := os.Getenv("PORT")
    if port == "" {
        port = "3001" // Default port
    }

    // HTTP server
    srv := &http.Server{
        Addr:    ":" + port,
        Handler: r,
    }

    // Start server
    go func() {
        log.Printf("Server is running on port %s", port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Failed to run server: %v", err)
        }
    }()

    // Wait for interrupt signal to gracefully shut down the server
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit

    log.Println("Shutting down server...")

    // Create a deadline to wait for ongoing requests
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Gracefully shutdown the server
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    log.Println("Server exiting")
}
