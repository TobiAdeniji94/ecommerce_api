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

### Example Endpoints:
- **POST /api/v1/users/register**: Register a new user.
- **POST /api/v1/users/login**: Authenticate a user and return a JWT token.
- **POST /api/v1/products**: Create a new product (Admin-only).
- **GET /api/v1/products**: Retrieve all products.
- **POST /api/v1/orders**: Place a new order.

### Endpoints

#### **POST /api/v1/users/register**
**Description**: Register a new user.  
**Request Body**:
```json
{
  "email": "user@example.com",
  "password": "strongpassword",
  "role": "user"
}
```
**Response (200)**:
```json
{
  "message": "User registered successfully",
  "data": {
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  }
}
```

---

#### **POST /api/v1/users/login**
**Description**: Authenticate a user and return a JWT token.  
**Request Body**:
```json
{
  "email": "user@example.com",
  "password": "strongpassword"
}
```
**Response (200)**:
```json
{
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  }
}
```

---

#### **POST /api/v1/products**
**Description**: Create a new product (Admin-only).  
**Headers**:
- `Authorization`: Bearer <JWT_TOKEN>
**Request Body**:
```json
{
  "name": "Wireless Mouse",
  "description": "Ergonomic wireless mouse with adjustable DPI",
  "price": 19.99,
  "stock": 100
}
```
**Response (200)**:
```json
{
  "message": "Product created successfully",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Wireless Mouse",
    "description": "Ergonomic wireless mouse with adjustable DPI",
    "price": 19.99,
    "stock": 100,
    "created_at": "2024-12-24T22:34:55Z",
    "updated_at": "2024-12-24T22:34:55Z"
  }
}
```

---

#### **GET /api/v1/products**
**Description**: Retrieve all products.  
**Headers**:
- `Authorization`: Bearer <JWT_TOKEN>
**Response (200)**:
```json
{
  "message": "Product(s) retrieved successfully",
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "Wireless Mouse",
      "description": "Ergonomic wireless mouse with adjustable DPI",
      "price": 19.99,
      "stock": 100,
      "created_at": "2024-12-24T22:34:55Z",
      "updated_at": "2024-12-24T22:34:55Z"
    }
  ]
}
```

---

#### **POST /api/v1/orders**
**Description**: Place a new order.  
**Headers**:
- `Authorization`: Bearer <JWT_TOKEN>
**Request Body**:
```json
{
  "items": [
    {
      "product_id": "123e4567-e89b-12d3-a456-426614174000",
      "quantity": 2
    }
  ]
}
```
**Response (200)**:
```json
{
  "message": "Order created successfully",
  "data": {
    "order_id": "123e4567-e89b-12d3-a456-426614174000",
    "status": "Pending",
    "items": [
      {
        "product_id": "123e4567-e89b-12d3-a456-426614174000",
        "quantity": 2
      }
    ]
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
```
