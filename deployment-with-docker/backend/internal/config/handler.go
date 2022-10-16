package config

import (
	albumHandler "postgres/internal/handler/album"
	albumUsecase "postgres/internal/usecase/album"
)

type Handler struct {
	AlbumHandler albumHandler.AlbumHandler
}

// Function to initialize handler
func InitHandler(albumUsecase albumUsecase.AlbumUsecase) Handler {
	return Handler{
		AlbumHandler: albumHandler.NewAlbumHandler(albumUsecase),
	}
}
