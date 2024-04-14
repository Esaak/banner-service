package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Esaak/banner-service/internal/banner/delivery/handlers"
	"github.com/Esaak/banner-service/internal/banner/repository/postgres"
	"github.com/Esaak/banner-service/internal/banner/usecase"
	"github.com/Esaak/banner-service/pkg/auth"
)

// Server represents the HTTP server
type Server struct {
	engine *gin.Engine
	port   int
}

// NewServer creates a new instance of Server
func NewServer(port int, db *gorm.DB, authService auth.AuthService) (*Server, error) {
	// Create repositories
	bannerRepo := postgres.NewBannerRepository(db)

	// Create use cases
	bannerUseCase := usecase.NewBannerUseCase(bannerRepo)

	// Create HTTP handlers
	bannerHandler := handlers.NewBannerHandler(bannerUseCase, authService)

	// Create Gin engine
	engine := gin.Default()

	// Register routes
	registerRoutes(engine, bannerHandler)

	return &Server{
		engine: engine,
		port:   port,
	}, nil
}

// Run starts the HTTP server
func (s *Server) Run() error {
	return s.engine.Run(fmt.Sprintf(":%d", s.port))
}

func registerRoutes(engine *gin.Engine, handler *handlers.BannerHandler) {
	// User routes
	engine.GET("/user_banner", handler.HandleGetUserBanner)

	// Admin routes
	admin := engine.Group("/admin")
	{
		admin.GET("/banners", handler.HandleGetBanners)
		admin.POST("/banners", handler.HandleCreateBanner)
		admin.PATCH("/banners/:id", handler.HandleUpdateBanner)
		admin.DELETE("/banners/:id", handler.HandleDeleteBanner)
	}
}
