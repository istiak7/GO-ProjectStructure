package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"CrudApp/delivery"
	"CrudApp/repository"
	"CrudApp/usecase"

	_ "github.com/lib/pq"
	_ "CrudApp/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

// @title Clean CRUD Postgres API
// @version 1.0
// @description Product Management API in Go
// @host localhost:8080
// @BasePath /

func runMigrations(db *sql.DB) {
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        log.Fatal("Migration Driver Error:", err)
    }

    // "file://migrations" looks for a folder named "migrations" in your current directory
    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations", 
        "postgres", 
        driver,
    )
    if err != nil {
        log.Fatal("Migration Sync Error:", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal("Migration Run Error:", err)
    }
    fmt.Println("Migrations applied!")
}


func main() {
	// 1. DATABASE CONNECTION
	// UPDATE user, password, and dbname with your Postgres credentials
	dsn := "postgres://postgres:istiakahmed@localhost:5432/gatekeeper?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
    log.Fatal("Database is unreachable:", err)
   } else {
    fmt.Println("Successfully connected to PostgreSQL!")
   }
     runMigrations(db)
	// 2. DEPENDENCY INJECTION
	repo := repository.NewPostgresRepo(db)
	uc := &usecase.ProductUsecase{Repo: repo}
	h := &delivery.ProductHandler{Usecase: uc}

	// 3. ROUTING (Using Go 1.22 Standard Library)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", h.GetProducts)
	mux.HandleFunc("POST /products", h.CreateProduct)
	mux.HandleFunc("PUT /products/{id}", h.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", h.DeleteProduct)
	
	// Swagger
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Server running at :8080")
	fmt.Println("Swagger UI: http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", mux))
}