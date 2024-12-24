// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a list of all orders placed by the authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get all orders for a user",
                "responses": {
                    "200": {
                        "description": "List of orders",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch orders",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows an authenticated user to place an order with one or more products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Place a new order",
                "parameters": [
                    {
                        "description": "Order payload",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PlaceOrderInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid order payload",
                        "schema": {
                            "$ref": "#/definitions/models.ValidationErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create order",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/orders/{id}/cancel": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows an authenticated user to cancel an order if it is in \"Pending\" status",
                "tags": [
                    "Orders"
                ],
                "summary": "Cancel an order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order canceled successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid order ID or status",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Order not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to cancel order",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/orders/{id}/status": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows an admin to update the status of an order",
                "tags": [
                    "Orders"
                ],
                "summary": "Update order status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update order status payload",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateOrderStatusInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order status updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid order ID or payload",
                        "schema": {
                            "$ref": "#/definitions/models.ValidationErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Order not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update order status",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves a list of all products available in the store",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "Product(s) retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to retrieve products",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows an admin user to add a new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "description": "Product payload",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid product payload",
                        "schema": {
                            "$ref": "#/definitions/models.ValidationErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create product",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves a single product by its ID",
                "tags": [
                    "Products"
                ],
                "summary": "Get product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid product ID",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Product not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to retrieve product",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows an admin user to update an existing product by ID",
                "tags": [
                    "Products"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated product payload",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product updated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid product ID or payload",
                        "schema": {
                            "$ref": "#/definitions/models.ValidationErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Product not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update product",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows an admin user to delete a product by ID",
                "tags": [
                    "Products"
                ],
                "summary": "Delete a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid product ID",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Product not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to delete product",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Authenticate a user with email and password, returning a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Authenticate a user",
                "parameters": [
                    {
                        "description": "User login payload",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Validation errors",
                        "schema": {
                            "$ref": "#/definitions/models.ValidationErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid email or password",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to generate token",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Create a new user in the system with email, password, and optional role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration payload",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User registered successfully",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Validation errors",
                        "schema": {
                            "$ref": "#/definitions/models.ValidationErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create user",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.LoginInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.OrderItemInput": {
            "type": "object",
            "properties": {
                "product_id binding:": {
                    "description": "Product ID",
                    "type": "string"
                },
                "quantity binding:": {
                    "description": "Quantity",
                    "type": "integer"
                }
            }
        },
        "models.PlaceOrderInput": {
            "type": "object",
            "properties": {
                "items": {
                    "description": "List of order items",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OrderItemInput"
                    }
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "models.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "models.UpdateOrderStatusInput": {
            "type": "object",
            "properties": {
                "status": {
                    "description": "order status",
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "models.ValidationError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ValidationErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ValidationError"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Use \"Bearer {your token}\" to authorize",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "E-Commerce API",
	Description:      "E-commerce API for managing orders and products.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
