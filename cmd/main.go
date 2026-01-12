package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"CrudApp/delivery"
	"CrudApp/repository"
	"CrudApp/usecase"

	// 1. Standard SQL and Migration imports
	_ "github.com/lib/pq"
	_ "CrudApp/docs"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	httpSwagger "github.com/swaggo/http-swagger"

	// 2. GORM imports
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func runMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Migration Driver Error:", err)
	}

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
	dsn := "postgres://postgres:istiakahmed@localhost:5432/gatekeeper?sslmode=disable"

	// A. Open the standard sql.DB first for Migrations
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Connection Error:", err)
	}

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	// B. Run Migrations using the standard sqlDB
	runMigrations(sqlDB)

	// C. Wrap the existing sqlDB into GORM
	gormDB, err := gorm.Open(gormPostgres.New(gormPostgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("GORM Initialization Error:", err)
	}

	// 2. DEPENDENCY INJECTION
	// Pass the gormDB to your repository
	repo := repository.NewPostgresRepo(gormDB) 
	uc := &usecase.ProductUsecase{Repo: repo}
	h := &delivery.ProductHandler{Usecase: uc}

	// 3. ROUTING
	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", h.GetProducts)
	mux.HandleFunc("POST /products", h.CreateProduct)
	mux.HandleFunc("PUT /products/{id}", h.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", h.DeleteProduct)

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Server running at :8080")
	fmt.Println("Swagger UI: http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", mux))
}