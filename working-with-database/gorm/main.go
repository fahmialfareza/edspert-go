package main

import (
	"fmt"
	"gorm/internal/config"
	"gorm/internal/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Get the config from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load db config
	db := config.OpenDB(os.Getenv("POSTGRES_URL"))
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

	// if err := handler.AlbumHandler.CreateAlbum(entity.Album{
	// 	Title:  "Pagi Yang Cerah",
	// 	Artist: "Peterpan",
	// 	Price:  50000,
	// }); err != nil {
	// 	panic(err)
	// }

	handler.AlbumHandler.BulkInsertAlbum([]entity.Album{
		{
			Title:  "Mandi Pagi",
			Artist: "Peterpan",
			Price:  50000,
		},
		{
			Title:  "Mandi Sore",
			Artist: "Peterpan",
			Price:  50000,
		},
	})

	albums, err := handler.AlbumHandler.GetAllAlbum()
	if err != nil {
		panic(err)
	}
	fmt.Println(albums)

	handler.AlbumHandler.UpdateAlbum(entity.Album{
		ID:     5,
		Title:  "Mandi Pagi2",
		Artist: "Peterpan",
		Price:  50000,
	})

	album, err := handler.AlbumHandler.GetAlbumById(5)
	if err != nil {
		panic(err)
	}
	fmt.Println(album)
}
