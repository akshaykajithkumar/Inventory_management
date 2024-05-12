# Inventory Management Platform Backend


## Objective

- This project aims to develop a robust backend system in Golang for an order and inventory management platform. The system features dynamic pricing for products influenced by factors such as demand and availability.
Framework

- The preferred framework for this project is Gin, a lightweight, fast, and robust middleware for API development in Golang.
Requirements
### 1. Database

Utilizes PostgreSQL as the Relational Database Management System (RDBMS).
Incorporates GORM as the Object-Relational Mapping (ORM) tool for seamless interaction with the database.
Design includes database models for tables such as products, orders, customers, inventory, etc., ensuring appropriate relationships between them.
SQL joins are utilized for data retrieval in APIs, and methodologies are implemented to optimize query response time.
Indexing is applied in the user table and order table for scalability and search performance.

 ### 2. APIs
#### User-Side APIs:

 Features include sign-up, login, product search, viewing products, placing orders, and accessing the user dashboard to view orders.

#### Admin-Side APIs:

Management of admin roles and inventory, including adding, removing, and updating products.
Order management with filters (user, product, etc.) and sorting options, including details of the ordering user.
APIs for generating statistics related to orders, users, inventory, etc., with appropriate filters and sorting capabilities.
Middleware is implemented for access control across different APIs.

### 3. Database Trigger for Dynamic Pricing

A database trigger is developed to adjust product pricing based on demand and availability.
A trigger is implemented to update the inventory when orders are placed.

## Overview

The backend system for order and inventory management implements dynamic pricing based on demand and availability. It adheres to the following key features:

- Clean Architecture for maintainability and scalability.
- PostgreSQL for efficient data storage.
- Gin Framework for API development.
- JWT for secure authentication and authorization.
- Dependency Injection and Compile-Time Dependency Injection for flexible component integration.
### Additional Features

Refresh Token and Access Token: Implemented secure authentication using JWT (JSON Web Tokens) for both access and refresh tokens. This ensures enhanced security and improved user experience by allowing seamless access to protected resources.

- Token Implementation Details

    Token Generation: Upon successful authentication, both an access token and a refresh token are generated and provided to the client.

    Access Token Expiry: The access token has a short expiration time to minimize the risk of unauthorized access.

    Refresh Token Expiry: Refresh tokens have a longer expiration time compared to access tokens, allowing users to obtain new access tokens without re-authentication.

    Token Refresh Mechanism: When the access token expires, the client uses the refresh token to request a new access token from the server.

    Revocation: Refresh tokens can be revoked on the server-side in case of security concerns or user logout, rendering them unusable for generating new access tokens.

### Indexing for Scalability and Search Performance

User Table Indexing: Indexes are created on user ID and username columns to accelerate user-related queries such as login, profile retrieval, and authentication.

Order Table Indexing: Indexes are established on order ID and user ID columns to optimize order-related operations such as tracking, history retrieval, and user-specific order listing.

Getting Started



```bash
Install and set up PostgreSQL.
    Create a dev.env file in the /pkg/config directory.
    Add PostgreSQL details to the dev.env file:
        DB_HOST: PostgreSQL host name.
        DB_NAME: PostgreSQL database name.
        DB_USER: PostgreSQL user name.
        DB_PORT: PostgreSQL port number.
        DB_PASSWORD: PostgreSQL password.
    Add a secret hash key to the dev.env file:
        KEY: your secret hash key.
    Run go run cmd/main.go in the terminal.
```
APIs
### User

- POST  /sign-up: Create a user account.
- GET  /log-in: User login.
- GET /product/search: Search products.
- GET /inventories/view/:id: View product details.
- POST /inventories/order: Place an order.
- GET /profile/orders: View past orders.

### Admin

-   GET /admin/log-in: Admin login.
-   GET /admin/users/list: View user list.
-  POST /admin/inventories/add: Add products.
-  DELETE /admin/inventories/delete: Remove products.
-  PUT /admin/inventories/update: Update product details.
-  GET /admin/orders/list: View order lists.
-  GET /admin/orders/:id :view specific order
-  GET /admin/orders/:id/status :changing the order status
-  GET /admin/order/stats: View order statistics.
-  GET /admin/user/stats: View user statistics.
-  GET /admin/inventory/stats: View inventory statistics.

 ### External Packages Used

  - [Gin](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
- [GORM](https://gorm.io/index.html) with [PostgresSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
- [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
- [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.
- [swag](https://github.com/swaggo/swag) converts Go annotations to Swagger Documentation 2.0 with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files)
