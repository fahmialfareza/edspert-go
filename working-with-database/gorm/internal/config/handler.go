package config

import (
	albumHandler "gorm/internal/handler/album"
	albumUsecase "gorm/internal/usecase/album"
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
