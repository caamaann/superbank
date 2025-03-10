package main

import (
	"log"

	"superbank/internal/config"
	"superbank/internal/handler"
	"superbank/internal/repository/postgres"
	"superbank/internal/service"
)

func main() {
	
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	
	db, err := postgres.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	
	customerRepo := postgres.NewCustomerRepository(db)
	userRepo := postgres.NewUserRepository(db)

	
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	customerService := service.NewCustomerService(customerRepo)
	userService := service.NewUserService(userRepo)

	
	server := handler.NewServer(cfg, authService, customerService, userService)
	server.Start()
}
