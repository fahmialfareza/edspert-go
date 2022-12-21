package main

import (
	"fmt"
	"log"
	"os"
	"postgres/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Get the environment
	env := os.Getenv("ENV")
	if env == "production" || env == "staging" {
		// Set to release mode
		gin.SetMode(gin.ReleaseMode)
	} else {
		// Get the config from .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// Initialize gin
	r := gin.Default()

	port := os.Getenv("PORT")

	// Load db config
	db, err := config.OpenDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// Load redis
	cache := config.OpenCache(os.Getenv("REDIS"), "")

	// Init clean arch
	repository := config.InitRepository(db, cache)
	usecase := config.InitUsecase(repository.AlbumRepository)
	handler := config.InitHandler(usecase.AlbumUsecase)

	// Create the API
	albumRoutes := r.Group("/api/v1/albums")
	{
		albumRoutes.GET("/", handler.AlbumHandler.GetAllAlbum)
		albumRoutes.POST("/", handler.AlbumHandler.Create)
		albumRoutes.POST("/batch", handler.AlbumHandler.BatchCreate)
		albumRoutes.GET("/:id", handler.AlbumHandler.Get)
		albumRoutes.PUT("/:id", handler.AlbumHandler.Update)
		albumRoutes.DELETE("/:id", handler.AlbumHandler.Delete)
	}

	// Run the gin gonic in port 5000
	runWithPort := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(runWithPort)
}
