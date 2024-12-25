# E-Commerce API

The **E-Commerce API** is a backend service designed to manage orders, products, and user authentication for an online store. It includes secure user registration and login features, product management, and order handling, with robust security and performance optimizations.

---

## Features

- **User Authentication**:
  - Register users with hashed passwords.
  - Login functionality with JWT-based authentication.
  - Role-based access control (e.g., admin-only routes).

- **Product Management**:
  - Create, read, update, and delete (CRUD) operations for products.
  - Admin-only access for creating, updating, and deleting products.

- **Order Management**:
  - Place and retrieve user orders.
  - Admin-only functionality for updating order status.

- **Swagger API Documentation**: 
  - Auto-generated and interactive documentation for easy API testing.

- **Rate Limiting**:
  - Per-client rate limiting to prevent abuse.

- **CORS Middleware**:
  - Protects the API from cross-origin requests while allowing specified domains.

---

## Table of Contents

- [Installation](#installation)
- [Setup](#setup)
- [API Documentation](#api-documentation)
- [Environment Variables](#environment-variables)
- [Database Schema](#database-schema)
- [How It Works](#how-it-works)
- [Testing](#testing)
- [License](#license)

---

## Installation

### Prerequisites

- **Go** (v1.18+)
- **PostgreSQL** (v14+)

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/TobiAdeniji94/ecommerce_api.git
   cd ecommerce_api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables by creating a `.env` file:
   ```bash
   touch .env
   ```

4. Populate the `.env` file (see [Environment Variables](#environment-variables)).

5. Run the project:
   ```bash
   go run main.go
   ```
   
6. Access the application:
   - API Base URL: [`https://ecommerce-api-vkui.onrender.com`](https://ecommerce-api-vkui.onrender.com)
   - Swagger Docs: [`https://ecommerce-api-vkui.onrender.com/swagger`](https://ecommerce-api-vkui.onrender.com/swagger/index.html)

---

## API Documentation

The Swagger UI is available at `/swagger` (e.g., [`https://ecommerce-api-vkui.onrender.com/swagger`](https://ecommerce-api-vkui.onrender.com/swagger/index.html)).

## **User Management**

### **Register a New User**
- **Method**: `POST`
- **Route**: `/api/v1/users/register`
- **Description**: Register a new user with email and password.
- **Access**: Public

#### **Request Payload**:
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "User registered successfully",
    "data": {
      "user_id": "uuid-1234-5678-91011"
    }
  }
  ```
- **Validation Error (400)**:
  ```json
  {
    "errors": [
      {
        "field": "email",
        "message": "Email is required"
      },
      {
        "field": "password",
        "message": "Password is required"
      }
    ]
  }
  ```

---

### **Login**
- **Method**: `POST`
- **Route**: `/api/v1/users/login`
- **Description**: Authenticate a user and issue a JWT for session management.
- **Access**: Public

#### **Request Payload**:
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Login successful",
    "data": {
      "token": "jwt-token",
      "user_id": "uuid-1234-5678-91011"
    }
  }
  ```
- **Invalid Credentials (401)**:
  ```json
  {
    "message": "Invalid email or password"
  }
  ```

---

## **Product Management** (Admin Privileges Required)

### **Create a Product**
- **Method**: `POST`
- **Route**: `/api/v1/products`
- **Description**: Create a new product.
- **Access**: Admin only
- **Headers**: `Authorization`: Bearer <JWT_TOKEN>

#### **Request Payload**:
```json
{
  "name": "Wireless Mouse",
  "description": "Ergonomic wireless mouse with adjustable DPI",
  "price": 19.99,
  "stock": 100
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Product created successfully",
    "data": {
      "id": "uuid-1234-5678-91011",
      "name": "Wireless Mouse",
      "description": "Ergonomic wireless mouse with adjustable DPI",
      "price": 19.99,
      "stock": 100,
      "created_at": "2024-12-25T10:00:00Z"
    }
  }
  ```

---

### **List All Products**
- **Method**: `GET`
- **Route**: `/api/v1/products`
- **Description**: Retrieve a list of all available products.
- **Access**: Authenticated users
- **Headers**: `Authorization`: Bearer <JWT_TOKEN>

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Products retrieved successfully",
    "data": [
      {
        "id": "uuid-1234-5678-91011",
        "name": "Wireless Mouse",
        "description": "Ergonomic wireless mouse with adjustable DPI",
        "price": 19.99,
        "stock": 100
      }
    ]
  }
  ```

---

### **Get a Product by ID**
- **Method**: `GET`
- **Route**: `/api/v1/products/{id}`
- **Description**: Retrieve details of a specific product by its ID.
- **Access**: Authenticated users
- **Headers**: `Authorization`: Bearer <JWT_TOKEN>

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Product retrieved successfully",
    "data": {
      "id": "uuid-1234-5678-91011",
      "name": "Wireless Mouse",
      "description": "Ergonomic wireless mouse with adjustable DPI",
      "price": 19.99,
      "stock": 100
    }
  }
  ```

---

### **Update a Product**
- **Method**: `PUT`
- **Route**: `/api/v1/products/{id}`
- **Description**: Update details of an existing product by its ID.
- **Access**: Admin only
- **Headers**: `Authorization`: Bearer <JWT_TOKEN>

#### **Request Payload**:
```json
{
  "name": "Updated Wireless Mouse",
  "description": "Updated ergonomic wireless mouse",
  "price": 24.99,
  "stock": 150
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Product updated successfully",
    "data": {
      "id": "uuid-1234-5678-91011",
      "name": "Updated Wireless Mouse",
      "description": "Updated ergonomic wireless mouse",
      "price": 24.99,
      "stock": 150
    }
  }
  ```

---

### **Delete a Product**
- **Method**: `DELETE`
- **Route**: `/api/v1/products/{id}`
- **Description**: Delete a product by its ID.
- **Access**: Admin only
- **Headers**: `Authorization`: Bearer <JWT_TOKEN>

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Product deleted successfully"
  }
  ```

---

## **Order Management**

### **Place an Order**
- **Method**: `POST`
- **Route**: `/api/v1/orders`
- **Description**: Place an order for one or more products.
- **Access**: Authenticated users
- **Headers**: `Authorization`: Bearer <JWT_TOKEN>

#### **Request Payload**:
```json
{
  "items": [
    {
      "product_id": "uuid-1234-5678-91011",
      "quantity": 2
    }
  ]
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Order placed successfully",
    "data": {
      "order_id": "uuid-1234-5678-91011",
      "status": "Pending",
      "items": [
        {
          "product_id": "uuid-1234-5678-91011",
          "quantity": 2
        }
      ]
    }
  }
  ```

---

### **List All Orders for a User**
- **Method**: `GET`
- **Route**: `/api/v1/orders`
- **Description**: List all orders placed by the authenticated user.
- **Access**: Authenticated users
- **Headers**: `Authorization`: Bearer <JWT_TOKEN>

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Orders retrieved successfully",
    "data": [
      {
        "order_id": "uuid-1234-5678-91011",
        "status": "Pending",
        "items": [
          {
            "product_id": "uuid-1234-5678-91011",
            "quantity": 2
          }
        ]
      }
    ]
  }
  ```

---

### **Cancel an Order**
- **Method**: `PUT`
- **Route**: `/api/v1/orders/{id}/cancel`
- **Description**: Cancel an order if it is still in the "Pending" status.
- **Access**: Authenticated users
- **Headers**: `Authorization`: Bearer <JWT_TOKEN>

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Order canceled successfully",
    "data": {
      "order_id": "uuid-1234-5678-91011",
      "status": "Canceled"
    }
  }
  ```

---

### **Update Order Status**
- **Method**: `PUT`
- **Route**: `/api/v1/orders/{id}/status`
- **Description**: Update the status of an order.
- **Access**: Admin only
- **Headers**: `Authorization`: Bearer <JWT_TOKEN>

#### **Request Payload**:
```json
{
  "status": "Shipped"
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Order status updated successfully",
    "data": {
      "order_id": "uuid-1234-5678-91011",
      "status": "Shipped"
    }
  }
  ```
---

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
JWT_SECRET=
PORT=3000
```

---

## **Database Schema**

### `users` Table

| Column       | Type       | Description             |
|--------------|------------|-------------------------|
| `id`         | UUID       | Primary key             |
| `email`      | VARCHAR(255) | Unique user email     |
| `password`   | VARCHAR(255) | Hashed password       |
| `role`       | VARCHAR(50)  | User role (default: user) |
| `created_at` | TIMESTAMP  | Timestamp of creation   |

### `products` Table

| Column       | Type       | Description                  |
|--------------|------------|------------------------------|
| `id`         | UUID       | Primary key                  |
| `name`       | VARCHAR(255) | Product name               |
| `description`| TEXT       | Product description          |
| `price`      | FLOAT      | Price of the product         |
| `stock`      | INT        | Available stock              |
| `created_at` | TIMESTAMP  | Timestamp of creation        |

### `orders` Table

| Column       | Type       | Description                  |
|--------------|------------|------------------------------|
| `id`         | UUID       | Primary key                  |
| `user_id`    | UUID       | Foreign key to `users` table |
| `status`     | VARCHAR(50)  | Order status (default: Pending) |
| `created_at` | TIMESTAMP  | Timestamp of creation        |

---

## How It Works

1. **User Authentication**:
   - Secure password hashing using bcrypt.
   - JWT-based authentication for secure API access.

2. **Product Management**:
   - Admins can create, update, and delete products.
   - Public access for viewing products.

3. **Order Management**:
   - Users can place orders and retrieve their order history.

4. **Rate Limiting**:
   - Implemented per-client to prevent API abuse.

5. **CORS Protection**:
   - Configured to allow only specific origins for API access.

---

## Testing

1. Use Postman or Swagger UI for manual testing of API endpoints.

---

## License

This project is licensed under the [MIT License](LICENSE).

---


Here’s the updated API documentation with JSON payloads and responses included for the described functionalities.

---

## **User Management**

### **Register a New User**
- **Method**: `POST`
- **Route**: `/api/v1/users/register`
- **Description**: Register a new user with email and password.
- **Access**: Public

#### **Request Payload**:
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "User registered successfully",
    "data": {
      "user_id": "uuid-1234-5678-91011"
    }
  }
  ```
- **Validation Error (400)**:
  ```json
  {
    "errors": [
      {
        "field": "email",
        "message": "Email is required"
      },
      {
        "field": "password",
        "message": "Password is required"
      }
    ]
  }
  ```

---

### **Login**
- **Method**: `POST`
- **Route**: `/api/v1/users/login`
- **Description**: Authenticate a user and issue a JWT for session management.
- **Access**: Public

#### **Request Payload**:
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Login successful",
    "data": {
      "token": "jwt-token",
      "user_id": "uuid-1234-5678-91011"
    }
  }
  ```
- **Invalid Credentials (401)**:
  ```json
  {
    "message": "Invalid email or password"
  }
  ```

---

## **Product Management** (Admin Privileges Required)

### **Create a Product**
- **Method**: `POST`
- **Route**: `/api/v1/products`
- **Description**: Create a new product.
- **Access**: Admin only

#### **Request Payload**:
```json
{
  "name": "Wireless Mouse",
  "description": "Ergonomic wireless mouse with adjustable DPI",
  "price": 19.99,
  "stock": 100
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Product created successfully",
    "data": {
      "id": "uuid-1234-5678-91011",
      "name": "Wireless Mouse",
      "description": "Ergonomic wireless mouse with adjustable DPI",
      "price": 19.99,
      "stock": 100,
      "created_at": "2024-12-25T10:00:00Z"
    }
  }
  ```

---

### **List All Products**
- **Method**: `GET`
- **Route**: `/api/v1/products`
- **Description**: Retrieve a list of all available products.
- **Access**: Authenticated users

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Products retrieved successfully",
    "data": [
      {
        "id": "uuid-1234-5678-91011",
        "name": "Wireless Mouse",
        "description": "Ergonomic wireless mouse with adjustable DPI",
        "price": 19.99,
        "stock": 100
      }
    ]
  }
  ```

---

### **Get a Product by ID**
- **Method**: `GET`
- **Route**: `/api/v1/products/{id}`
- **Description**: Retrieve details of a specific product by its ID.
- **Access**: Authenticated users

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Product retrieved successfully",
    "data": {
      "id": "uuid-1234-5678-91011",
      "name": "Wireless Mouse",
      "description": "Ergonomic wireless mouse with adjustable DPI",
      "price": 19.99,
      "stock": 100
    }
  }
  ```

---

### **Update a Product**
- **Method**: `PUT`
- **Route**: `/api/v1/products/{id}`
- **Description**: Update details of an existing product by its ID.
- **Access**: Admin only

#### **Request Payload**:
```json
{
  "name": "Updated Wireless Mouse",
  "description": "Updated ergonomic wireless mouse",
  "price": 24.99,
  "stock": 150
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Product updated successfully",
    "data": {
      "id": "uuid-1234-5678-91011",
      "name": "Updated Wireless Mouse",
      "description": "Updated ergonomic wireless mouse",
      "price": 24.99,
      "stock": 150
    }
  }
  ```

---

### **Delete a Product**
- **Method**: `DELETE`
- **Route**: `/api/v1/products/{id}`
- **Description**: Delete a product by its ID.
- **Access**: Admin only

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Product deleted successfully"
  }
  ```

---

## **Order Management**

### **Place an Order**
- **Method**: `POST`
- **Route**: `/api/v1/orders`
- **Description**: Place an order for one or more products.
- **Access**: Authenticated users

#### **Request Payload**:
```json
{
  "items": [
    {
      "product_id": "uuid-1234-5678-91011",
      "quantity": 2
    }
  ]
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Order placed successfully",
    "data": {
      "order_id": "uuid-1234-5678-91011",
      "status": "Pending",
      "items": [
        {
          "product_id": "uuid-1234-5678-91011",
          "quantity": 2
        }
      ]
    }
  }
  ```

---

### **List All Orders for a User**
- **Method**: `GET`
- **Route**: `/api/v1/orders`
- **Description**: List all orders placed by the authenticated user.
- **Access**: Authenticated users

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Orders retrieved successfully",
    "data": [
      {
        "order_id": "uuid-1234-5678-91011",
        "status": "Pending",
        "items": [
          {
            "product_id": "uuid-1234-5678-91011",
            "quantity": 2
          }
        ]
      }
    ]
  }
  ```

---

### **Cancel an Order**
- **Method**: `PUT`
- **Route**: `/api/v1/orders/{id}/cancel`
- **Description**: Cancel an order if it is still in the "Pending" status.
- **Access**: Authenticated users

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Order canceled successfully",
    "data": {
      "order_id": "uuid-1234-5678-91011",
      "status": "Canceled"
    }
  }
  ```

---

### **Update Order Status**
- **Method**: `PUT`
- **Route**: `/api/v1/orders/{id}/status`
- **Description**: Update the status of an order.
- **Access**: Admin only

#### **Request Payload**:
```json
{
  "status": "Shipped"
}
```

#### **Response**:
- **Success (200)**:
  ```json
  {
    "message": "Order status updated successfully",
    "data": {
      "order_id": "uuid-1234-5678-91011",
      "status": "Shipped"
    }
  }
  ```

--- 

### **Notes**
- Ensure proper error handling for invalid inputs and missing parameters.
- All requests requiring authentication must include the `Authorization` header with the format:
  ```json
  {
    "Authorization": "Bearer <JWT_TOKEN>"
  }
  ```
