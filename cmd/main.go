package main

import (
	"database/sql"
	"log"

	"github.com/Esaak/banner-service/configs"
	"github.com/Esaak/banner-service/internal/server"
	"github.com/Esaak/banner-service/pkg/auth"
	"github.com/Esaak/banner-service/pkg/database"
	"github.com/Esaak/banner-service/pkg/logger"
)

func main() {
	// Load configuration
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	sqlDB, err := sql.Open("postgres", config.PostgresURL)
	if err := sqlDB.Ping(); err != nil {
		log.Println("DB Ping error: ", err.Error())
		return
	}
	defer sqlDB.Close()

	db, err := database.NewDB(sqlDB)

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Create auth service
	authService := auth.NewAuthService(config.UserSecret, config.AdminSecret)

	// Create HTTP server
	srv, err := server.NewServer(config.Port, db, authService)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	logger.Info("Server started on port", config.Port)

	// Start HTTP server
	if err := srv.Run(); err != nil {
		logger.Error(err)
	}
}
