# CrudApp

A RESTful CRUD API built with Go, using SQLx for database operations, PostgreSQL for data storage, and golang-migrate for schema management.

## Features

- Clean Architecture (Handler → UseCase → Repository)
- SQLx for efficient SQL queries
- PostgreSQL database
- Automatic database migrations
- Input validation with DTOs
- CORS support for frontend integration
- Structured JSON responses

## Project Structure

```
CrudApp/
├── cmd
|   └──main.go                   # Application entry point
├── common                       # Utilities & middleware
|   ├──middleware
|      └──cors.go
├──utilities
|  └──response.go  
|  └──validator.go           
├── config                       # Database & dependency configuration
|   └──app.go 
|   └──database.go 
|   └──dependencies.go            
├── delivery                     # HTTP handlers & DTOs
|   ├──Product
|       ├──dtos
|          └──CreateProductDto.go
|          └──UpdateProductDto.go
|      └──ProductHandler.go
|   ├──Supplier
|       ├──dtos
|          └──CreateSupplierDto.go
|       └──SupplierHandler.go
├── docs                         # API documentation
├── domain                       # Business entities & interfaces
|   └──Product.go
|   └──Supplier.go
├── migrations                   # Database migration files
├── repository                   # Database operations
|   └──ProductRepository.go
|   └──SupplierRepository.go
├── route                        # Route definitions
|   └──Routes.go
|   └──SupplierRoutes.go
|   └──SwaggerRoutes.go
├── usecase                      # Business logic
|   └──ProductUseCase.go
|   └──SupplierUseCase.go
├── go.mod
├── go.sum
   
```

## Prerequisites

- Go 1.25+
- PostgreSQL
- golang-migrate CLI

## Installation

### 1. Initialize Project

```bash
go mod init CrudApp
```

### 2. Install Dependencies

```bash
go get github.com/jmoiron/sqlx \
       github.com/lib/pq \
       github.com/golang-migrate/migrate/v4 \
       github.com/go-playground/validator/v10
```

### 3. Install Migration CLI

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Database Setup

### Configure Connection

Update `config/database.go` with your PostgreSQL credentials:

```go
dsn := "postgres://user:password@localhost:5432/dbname?sslmode=disable"
```

### Create Migration

```bash
migrate create -ext sql -dir migrations -seq create_products_table
```

### Run Migrations Manually

```bash
migrate -path migrations \
-database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" \
up
```

### Automatic Migrations

The application runs migrations automatically on startup via `config/database.go`.

## Running the Application

```bash
go run cmd/main.go
```

Server starts at `http://localhost:8080`

## API Endpoints

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/CreateProducts` | Create a new product |
| GET | `/GetProducts/{id}` | Get product by ID |
| PUT | `/UpdateProducts/{id}` | Update product |
| DELETE | `/DeleteProducts/{id}` | Delete product |

### Request/Response Format

**Create Product:**

```json
POST /CreateProducts
{
  "name": "Gaming Keyboard",
  "price": 75.50
}
```

**Response:**

```json
{
  "success": true,
  "message": "Product created successfully",
  "data": {
    "id": 1,
    "name": "Gaming Keyboard",
    "price": 75.50,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### Data Flow

```
HTTP Request → Handler → UseCase → Repository → Database
```

### Key Components

**Validation:**
- Input validation via `go-playground/validator`
- Automatic error responses for invalid data

## CORS Configuration

CORS is enabled in `common/middleware/cors.go` to allow frontend connections:

```go
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
```

