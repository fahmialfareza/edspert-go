package main

import (
	"fmt"
	"gorm/internal/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// load db config
	db := config.OpenDB(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// init clean arch
	repository := config.InitRepository(db)
	usecase := config.InitUsecase(repository.AlbumRepository)
	handler := config.InitHandler(usecase.AlbumUsecase)

	// using handler
	// Get all albums
	albums, err := handler.AlbumHandler.GetAllAlbum()
	if err != nil {
		panic(err)
	}
	fmt.Println(albums)
}
