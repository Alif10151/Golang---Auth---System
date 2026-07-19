# Golang Auth System 🚀

A secure and scalable authentication system built with **Golang**.
This project provides user authentication, authorization, and user management functionality using industry-standard backend practices.

The application focuses on clean API design, secure password handling, JWT-based authentication, middleware implementation, and database integration.

---

## 📌 Overview

Golang Auth System is a RESTful API backend that allows users to:

* Create an account
* Authenticate securely
* Access protected resources
* Manage user information
* Perform user-related CRUD operations

The project is designed with modular package organization to maintain clean and maintainable backend code.

---

# ✨ Features

## Authentication & Authorization

* User Registration
* User Login
* Secure password hashing using bcrypt
* JWT-based authentication
* Protected API routes
* User profile access
* Logout functionality

## User Management

* Retrieve all users
* Retrieve user by ID
* Update user information
* Delete user records

## Backend Engineering

* RESTful API design
* PostgreSQL database integration
* GORM ORM implementation
* Environment-based configuration
* Request validation
* Standardized API responses
* Middleware-based architecture
* Request logging
* CORS handling

---

# 🛠️ Technology Stack

| Technology | Purpose               |
| ---------- | --------------------- |
| Go         | Backend Development   |
| net/http   | HTTP Server & Routing |
| PostgreSQL | Database              |
| GORM       | Database ORM          |
| JWT        | Authentication        |
| bcrypt     | Password Security     |
| Postman    | API Testing           |

---

# 🏗️ Project Architecture

The project follows a modular backend structure:

```text
Client
  |
  ↓
Middleware Layer
  |
  ↓
Handler Layer
  |
  ↓
Database Layer
  |
  ↓
PostgreSQL
```

### Package Responsibilities

**Handlers**

* Handles HTTP requests and responses
* Manages API endpoints

**Middleware**

* Authentication checks
* Request logging
* CORS handling

**Models**

* Defines database entities

**Database**

* Manages database connection and configuration

**Utils**

* Provides reusable helper functions

---

# 📂 Project Structure

```text
Golang-Auth-System

├── cmd
│   └── main.go
│
├── internal
│   │
│   ├── db
│   │   └── db.go
│   │
│   ├── handlers
│   │   ├── auth.go
│   │   ├── logout_handler.go
│   │   ├── profile_handler.go
│   │   └── user.go
│   │
│   ├── middleware
│   │   ├── auth.go
│   │   ├── cors.go
│   │   └── logger.go
│   │
│   ├── models
│   │   ├── response.go
│   │   └── user.go
│   │
│   └── utils
│       ├── jwt.go
│       └── response.go
│
├── go.mod
├── go.sum
└── README.md
```

---

# ⚙️ Setup Instructions

## Clone the Repository

```bash
git clone <https://github.com/Alif10151/Golang---Auth---System >

cd Golang-Auth-System
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Environment Configuration

Create your own environment configuration file 

> Sensitive credentials should never be committed to version control.

---

## Run the Application

```bash
go run ./cmd
```

The server will start successfully after completing the required configuration.

---

# 🔐 Security Practices

This project follows backend security practices:

* Passwords are stored using secure hashing algorithms
* Authentication is handled through signed JWT tokens
* Sensitive configuration values are managed through environment variables
* Protected routes require valid authentication
* Database operations are handled securely through ORM methods

---

# 📡 API Modules

## Authentication

| Method | Endpoint    | Description        |
| ------ | ----------- | ------------------ |
| POST   | `/register` | Create new account |
| POST   | `/login`    | Authenticate user  |
| POST   | `/logout`   | Logout user        |

---

## User Management

| Method | Endpoint       | Description                    |
| ------ | -------------- | ------------------------------ |
| GET    | `/profile`     | Get authenticated user profile |
| GET    | `/users`       | Retrieve users                 |
| GET    | `/users/{id}`  | Retrieve user by ID            |
| PUT    | `/update_user` | Update user information        |
| DELETE | `/delete_user` | Delete user                    |

---

# 🧪 Testing

API endpoints were tested using:

* Postman
* Manual API testing
* Different authentication scenarios

---

# 🚀 Future Improvements

Possible improvements:

* Refresh token mechanism
* API documentation using Swagger
* Docker containerization
* Deployment pipeline
* Rate limiting

---

# Author

**Alif**
Backend Developer (Golang) | Golang Enthusiast

---

## License

This project is created for learning, development, and portfolio purposes.

