package main

import (
	"healtech-backend/server/internal/config"
	"healtech-backend/server/internal/handler"
	"healtech-backend/server/internal/middleware"
	"healtech-backend/server/internal/repository"
	"healtech-backend/server/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()
	repository.InitDB(cfg)
	authService := service.NewAuthService(cfg)
	docService := service.NewDocumentService(cfg)

	// 4. Initialize Handlers
	h := handler.NewHandler(docService, authService)

	// 5. Setup Router
	r := gin.Default()

	// Public Routes
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(authService))
	{
		protected.GET("/documents", h.ListDocuments)
		protected.POST("/documents/upload", h.UploadDocument)
		protected.GET("/documents/:id/download", h.DownloadDocument)
		protected.DELETE("/documents/:id", h.DeleteDocument)
	}
	log.Printf("Starting server on port: %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
