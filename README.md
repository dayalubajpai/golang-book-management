# Bookstore API

A Go-based RESTful API for managing a bookstore, built with Gin framework and MySQL database.

## 🚀 Features

- RESTful API endpoints for book management
- MySQL database integration
- Docker containerization
- PHPMyAdmin for database management
- Secure environment configuration

## 🛠️ Technologies

- Go 1.23.3
- Gin Web Framework
- MySQL (MariaDB 10.5)
- Docker & Docker Compose
- PHPMyAdmin

## 📋 Prerequisites

- Docker and Docker Compose
- Go 1.23.3 or later
- Git

## 🔧 Installation & Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd bookstore
   ```

2. Start the Docker containers:
   ```bash
   docker-compose up -d
   ```

3. Verify the containers are running:
   ```bash
   docker ps
   ```

4. Build and run the application:
   ```bash
   go mod download
   go run cmd/main.go
   ```

## 🔌 Database Configuration

The database is automatically configured with Docker Compose with the following credentials:

- **Database Name:** bookstore
- **Username:** admin
- **Password:** adminpassword
- **Port:** 3306

PHPMyAdmin is available at:
- **URL:** http://localhost:8080
- **Username:** admin
- **Password:** adminpassword

## 🏗️ Project Structure

bookstore/

├── cmd/ # Application entrypoint

├── pkg/ # Package code

├── docker-compose.yml

├── go.mod

├── go.sum

└── README.md

## Thank you!!
