package config

import albumHandler "gorm/internal/handler/album"

type Handler struct {
	AlbumHandler albumHandler.AlbumHandler
}

func InitHandler(albumHandler albumHandler.AlbumHandler) Handler {
	return Handler{
		AlbumHandler: albumHandler,
	}
}