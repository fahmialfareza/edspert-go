package config

import albumHandler "postgres/internal/handler/album"

type Handler struct {
	AlbumHandler albumHandler.AlbumHandler
}

// Function to initialize handler
func InitHandler(albumHandler albumHandler.AlbumHandler) Handler {
	return Handler{
		AlbumHandler: albumHandler,
	}
}
