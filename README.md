# Task Management System API

A scalable backend application built using GoLang, Gin Framework, MongoDB, JWT Authentication, and Role-Based Access Control (RBAC).

This project includes:
- User Authentication
- JWT Authorization
- Role-Based Access
- CRUD APIs for Tasks
- Admin APIs
- Docker Support
- MongoDB Integration
- Input Validation
- Secure Password Hashing

---

# Tech Stack

## Backend
- GoLang
- Gin Framework
- MongoDB
- JWT
- bcrypt
- Docker

## Frontend
- React.js

---

# Features

## Authentication
- User Registration
- User Login
- JWT Token Authentication
- Password Hashing using bcrypt

## Role-Based Access
- Admin
- User

## Task Management
- Create Task
- Get Tasks
- Update Task
- Delete Task

## Admin Features
- Get All Users
- Delete Users

## Security Features
- JWT Authentication
- Password Hashing
- Protected Routes
- Input Validation
- Environment Variables

---

# Folder Structure

```bash
backend/
│
├── cmd/
│   └── main.go
│
├── config/
│   └── db.go
│
├── controllers/
│   ├── authController.go
│   ├── taskController.go
│   └── adminController.go
│
├── middleware/
│   ├── authMiddleware.go
│   └── roleMiddleware.go
│
├── models/
│   ├── user.go
│   └── task.go
│
├── routes/
│   ├── authRoutes.go
│   ├── taskRoutes.go
│   └── adminRoutes.go
│
├── utils/
│   ├── hash.go
│   ├── jwt.go
│   └── response.go
│
├── validators/
│   └── validation.go
│
├── Dockerfile
├── docker-compose.yml
├── .env
├── go.mod
└── README.md
```

---

# API Endpoints

# Authentication APIs

## Register User

```http
POST /api/v1/auth/register
```

### Request Body

```json
{
  "name": "Anjali",
  "email": "anjali@gmail.com",
  "password": "123456",
  "role": "user"
}
```

---

## Login User

```http
POST /api/v1/auth/login
```

### Request Body

```json
{
  "email": "anjali@gmail.com",
  "password": "123456"
}
```

---

# Task APIs

## Create Task

```http
POST /api/v1/tasks/
```

### Headers

```http
Authorization: Bearer TOKEN
```

### Request Body

```json
{
  "title": "Build Backend",
  "description": "Complete Go Assignment",
  "status": "pending"
}
```

---

## Get Tasks

```http
GET /api/v1/tasks/
```

---

## Update Task

```http
PUT /api/v1/tasks/:id
```

---

## Delete Task

```http
DELETE /api/v1/tasks/:id
```

---

# Admin APIs

## Get All Users

```http
GET /api/v1/admin/users
```

---

## Delete User

```http
DELETE /api/v1/admin/users/:id
```

---

# Environment Variables

Create a `.env` file inside backend folder.

```env
PORT=8080

MONGO_URI=mongodb://mongodb:27017

JWT_SECRET=mysecretkey
```

---

# Local Setup

# Clone Repository

```bash
git clone YOUR_GITHUB_REPO
```

---

# Install Dependencies

```bash
go mod tidy
```

---

# Run Project

```bash
go run cmd/main.go
```

---

# Docker Setup

## Build & Run

```bash
docker-compose up --build
```

---

## Stop Containers

```bash
docker-compose down
```

---

# Scalability Notes

## 1. Microservices Architecture
Authentication, tasks, and notifications can be separated into independent services.

## 2. Redis Caching
Redis can be integrated to cache frequently accessed APIs and reduce database load.

## 3. Load Balancing
Multiple backend instances can run behind a load balancer for high availability.

## 4. Queue Systems
Kafka or RabbitMQ can be used for asynchronous processing and notifications.

## 5. Docker Deployment
Containerized deployment improves scalability and portability.

---

# Security Practices

- JWT Authentication
- Password Hashing using bcrypt
- Protected Routes
- Input Validation
- Environment Variables
- Role-Based Access Control

---

# Future Improvements

- Swagger API Documentation
- Refresh Tokens
- Email Verification
- Pagination
- Search & Filters
- Rate Limiting
- Redis Caching
- Kubernetes Deployment

---

# Author

Anjali
```