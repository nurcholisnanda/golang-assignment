package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nurcholisnanda/golang-assignment/internal/repository"
	"github.com/nurcholisnanda/golang-assignment/internal/service"
	"gorm.io/gorm"
)

// Setup User router group
func AddRoutes(g *gin.Engine, db *gorm.DB) {
	// Initialize repository and service
	userRepo := repository.NewUserRepoImpl(db)
	recordService := service.NewService(userRepo)

	// Create a new handler with the service
	handler := NewHandler(recordService)

	// Define routes
	g.GET("/records", handler.GetRecords)
}
