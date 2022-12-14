package config

import (
	albumRepository "postgres/internal/repository/album"
	albumUsecase "postgres/internal/usecase/album"
)

type Usecase struct {
	AlbumUsecase albumUsecase.AlbumUsecase
}

// Function to initialize usecase
func InitUsecase(albumRepository albumRepository.AlbumRepository) Usecase {
	return Usecase{
		AlbumUsecase: albumUsecase.NewAlbumUsecase(albumRepository),
	}
}
