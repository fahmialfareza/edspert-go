package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"postgres/internal/config"
	"postgres/internal/entity"

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
	// album, err := handler.AlbumHandler.Create(&entity.Album{
	// 	Title:  "Apa Ini 2",
	// 	Artist: "Peterpan",
	// 	Price:  50000,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// b, _ := json.Marshal(album)
	// fmt.Println(string(b))
	// fmt.Println("Berhasil")

	// Batch create
	// albums, err := handler.AlbumHandler.BatchCreate([]entity.Album{
	// 	{
	// 		Title:  "e",
	// 		Artist: "e",
	// 		Price:  12345,
	// 	},
	// 	{
	// 		Title:  "f",
	// 		Artist: "f",
	// 		Price:  67890,
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// b, _ := json.Marshal(albums)
	// fmt.Println(string(b))
	// fmt.Println("Berhasil")

	// Update
	album, err := handler.AlbumHandler.Update(entity.Album{
		ID:     9,
		Title:  "FFFF",
		Artist: "FFFF",
		Price:  10000,
	})
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(album)
	fmt.Println(string(b))
	fmt.Println("Berhasil")
}
