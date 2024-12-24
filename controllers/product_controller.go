package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/TobiAdeniji94/ecommerce_api/config"
	"github.com/TobiAdeniji94/ecommerce_api/models"
)

// CreateProduct allows an admin user to add a new product
// CreateProduct godoc
// @Summary Create a new product
// @Description Allows an admin user to add a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.ProductInput true "Product payload"
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse "Product created successfully"
// @Failure 400 {object} models.ValidationErrorResponse "Invalid product payload"
// @Failure 500 {object} models.ErrorResponse "Failed to create product"
// @Router /products [post]
func CreateProduct(c *gin.Context) {
    var input models.ProductInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, models.ValidationErrorResponse{
            Errors: []models.ValidationError{
                {Field: "payload", Message: err.Error()},
            },
        })
        return
    }

    // Map the input to the Product model
    product := models.Product{
        Name:        input.Name,
        Description: input.Description,
        Price:       input.Price,
        Stock:       input.Stock,
    }

    // Insert the Product model into the database
    if err := config.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to create product"})
        return
    }

    c.JSON(http.StatusOK, models.SuccessResponse{
        Message: "Product created successfully",
        Data:    product,
    })
}

// GetProducts lists all products
// GetProducts godoc
// @Summary Get all products
// @Description Retrieves a list of all products available in the store
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse "Product(s) retrieved successfully"
// @Failure 500 {object} models.ErrorResponse "Failed to retrieve products"
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Product(s) retrieved successfully",
		Data:    products,
	})
}

// GetProductByID retrieves a single product by ID
// GetProductByID godoc
// @Summary Get product by ID
// @Description Retrieves a single product by its ID
// @Tags Products
// @Param id path string true "Product ID"
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse "Product retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid product ID"
// @Failure 404 {object} models.ErrorResponse "Product not found"
// @Failure 500 {object} models.ErrorResponse "Failed to retrieve product"
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid product ID"})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Product not found"})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Product retrieved successfully",
		Data:    product,
	})
}

// UpdateProduct modifies an existing product (admin only)
// UpdateProduct godoc
// @Summary Update a product
// @Description Allows an admin user to update an existing product by ID
// @Tags Products
// @Param id path string true "Product ID"
// @Param product body models.ProductInput true "Updated product payload"
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse "Product updated successfully"
// @Failure 400 {object} models.ValidationErrorResponse "Invalid product ID or payload"
// @Failure 404 {object} models.ErrorResponse "Product not found"
// @Failure 500 {object} models.ErrorResponse "Failed to update product"
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid product ID"})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Product not found"})
		return
	}

	var updateInput models.ProductInput
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, models.ValidationErrorResponse{
			Errors: []models.ValidationError{
				{Field: "payload", Message: err.Error()},
			},
		})
		return
	}

	// Update fields
	product.Name = updateInput.Name
	product.Description = updateInput.Description
	product.Price = updateInput.Price
	product.Stock = updateInput.Stock

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Product updated successfully",
		Data:    product,
	})
}

// DeleteProduct deletes a product by ID (admin only)
// DeleteProduct godoc
// @Summary Delete a product
// @Description Allows an admin user to delete a product by ID
// @Tags Products
// @Param id path string true "Product ID"
// @Security BearerAuth
// @Success 200 {object} models.SuccessResponse "Product deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid product ID"
// @Failure 404 {object} models.ErrorResponse "Product not found"
// @Failure 500 {object} models.ErrorResponse "Failed to delete product"
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid product ID"})
		return
	}

	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Product deleted successfully",
	})
}
