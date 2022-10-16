package main

import (
	"fmt"
	"log"
	"os"
	"postgres/internal/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Get the config from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load db config
	db, err := config.OpenDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// Init clean arch
	repository := config.InitRepository(db)
	usecase := config.InitUsecase(repository.AlbumRepository)
	handler := config.InitHandler(usecase.AlbumUsecase)

	// Using the handler
	// You can use/change the code below if you want to test the handler
	albums, err := handler.AlbumHandler.GetAllAlbum()
	if err != nil {
		panic(err)
	}
	fmt.Println(albums)
}
