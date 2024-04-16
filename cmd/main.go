package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nurcholisnanda/golang-assignment/configs/database"
	"github.com/nurcholisnanda/golang-assignment/internal/api"
)

func main() {
	//setup database
	db, err := database.NewDatabase()
	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate() // Automatically migrate database schema
	gormDB := db.GetDB()

	r := gin.Default() // Create a new Gin router with default middleware
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"message": "Golang-Assignment",
		})
	})

	// Setup API routes
	api.AddRoutes(r, gormDB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if environment variable is not set
	}

	if err = r.Run(":" + port); err != nil {
		log.Panic(err)
	}
}
