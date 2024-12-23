package routes

import (
    "github.com/gin-gonic/gin"
    
    "github.com/TobiAdeniji94/ecommerce_api/controllers"
    "github.com/TobiAdeniji94/ecommerce_api/middleware"
)

func InitializeRoutes(r *gin.Engine) {

    // Public routes: user registration/login
    userGroup := r.Group("/users")
    {
        userGroup.POST("/register", controllers.RegisterUser)
        userGroup.POST("/login", controllers.LoginUser)
    }

    // Authenticated routes
    protected := r.Group("/")
    protected.Use(middleware.AuthMiddleware) // Checks JWT

    // Product routes (admin only for create/update/delete)
    productGroup := protected.Group("/products")
    {
        productGroup.POST("", middleware.AdminMiddleware, controllers.CreateProduct)
        productGroup.GET("", controllers.GetProducts)
        productGroup.GET("/:id", controllers.GetProductByID)
        productGroup.PUT("/:id", middleware.AdminMiddleware, controllers.UpdateProduct)
        productGroup.DELETE("/:id", middleware.AdminMiddleware, controllers.DeleteProduct)
    }

    // Order routes
    orderGroup := protected.Group("/orders")
    {
        orderGroup.POST("", controllers.PlaceOrder)
        orderGroup.GET("", controllers.GetUserOrders)
        orderGroup.PUT("/:id/cancel", controllers.CancelOrder)
        orderGroup.PUT("/:id/status", middleware.AdminMiddleware, controllers.UpdateOrderStatus)
    }
}
