package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"photographer/internal/config"
	"photographer/internal/repository"
	"photographer/internal/service"
	http_handler "photographer/internal/transport/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "photographer/docs" // импорт сгенерированных документов
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgresConfig.User,
		cfg.PostgresConfig.Password,
		cfg.PostgresConfig.Host,
		cfg.PostgresConfig.Port,
		cfg.PostgresConfig.DBName,
		cfg.PostgresConfig.SSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	if err = runMigrations(connStr); err != nil {
		log.Fatalln("migration error:", err)
	}

	repo := repository.New(db)
	_service := service.New(repo)
	router := http_handler.NewHandler(_service)

	// Добавляем маршрут для Swagger UI
	muxRouter := router.Handle()
	muxRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server starting on port 8080...")
	if err = http.ListenAndServe(":8080", muxRouter); err != nil {
		log.Fatal(err)
	}
}

func runMigrations(dbURL string) error {
	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		return fmt.Errorf("failed to init migrations: %w", err)
	}

	if err = m.Up(); !errors.Is(err, migrate.ErrNoChange) && err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
