package routes

import (
	"net/http"
    "github.com/gin-gonic/gin"

    "github.com/TobiAdeniji94/ecommerce_api/controllers"
    "github.com/TobiAdeniji94/ecommerce_api/middleware"
)

func InitializeRoutes(r *gin.Engine) {

	// Welcome message
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the E-commerce API",
		})
	})

    // API Versioning
    api := r.Group("/api/v1")
    {
        // Public Routes: User registration and login
        userGroup := api.Group("/users")
        {
            userGroup.POST("/register", controllers.RegisterUser)
            userGroup.POST("/login", controllers.LoginUser)
        }

        // Protected Routes: Requires Authentication
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware) // JWT authentication middleware

        // Product Routes: Admin-only for create, update, delete
        productGroup := protected.Group("/products")
        {
            productGroup.POST("", middleware.AdminMiddleware, controllers.CreateProduct)  // Create a product
            productGroup.GET("", controllers.GetProducts)                                // List all products
            productGroup.GET("/:id", controllers.GetProductByID)                         // Get product by ID
            productGroup.PUT("/:id", middleware.AdminMiddleware, controllers.UpdateProduct) // Update a product
            productGroup.DELETE("/:id", middleware.AdminMiddleware, controllers.DeleteProduct) // Delete a product
        }

        // Order Routes: Authenticated users and admin access
        orderGroup := protected.Group("/orders")
        {
            orderGroup.POST("", controllers.PlaceOrder)                          // Place a new order
            orderGroup.GET("", controllers.GetUserOrders)                        // List user orders
            orderGroup.PUT("/:id/cancel", controllers.CancelOrder)               // Cancel an order
            orderGroup.PUT("/:id/status", middleware.AdminMiddleware, controllers.UpdateOrderStatus) // Update order status (Admin)
        }
    }
}
