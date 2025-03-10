package handler

import (
	"fmt"
	"log"
	"time"

	"superbank/internal/config"
	"superbank/internal/middleware"
	"superbank/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router          *gin.Engine
	config          *config.Config
	authService     service.AuthService
	customerService service.CustomerService
	userService     service.UserService
}

func NewServer(
	config *config.Config,
	authService service.AuthService,
	customerService service.CustomerService,
	userService service.UserService, 
) *Server {
	server := &Server{
		config:          config,
		authService:     authService,
		customerService: customerService,
		userService:     userService, 
	}

	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	router := gin.Default()

	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     s.config.CorsOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	
	router.POST("/api/auth/login", s.handleLogin)
	router.POST("/api/auth/register", s.handleAddUser)

	
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(s.authService))
	{
		protected.GET("/customers/search", s.handleCustomerSearch)
		protected.GET("/customers/:id", s.handleGetCustomer)
	}

	s.router = router
}

func (s *Server) Start() {
	log.Printf("Server running on port %s", s.config.Port)
	if err := s.router.Run(fmt.Sprintf(":%s", s.config.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
