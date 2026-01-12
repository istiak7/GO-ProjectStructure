package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "postgres://postgres:istiakahmed@localhost:5432/gatekeeper?sslmode=disable"

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Connection Error:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	runMigrations(sqlDB)

	gormDB, err := gorm.Open(gormPostgres.New(gormPostgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("GORM Initialization Error:", err)
	}

	return gormDB
}

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
