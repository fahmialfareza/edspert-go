package config

import albumUsecase "postgres/internal/usecase/album"

type Usecase struct {
	AlbumUsecase albumUsecase.AlbumUsecase
}

// Function to initialize usecase
func InitUsecase(albumUsecase albumUsecase.AlbumUsecase) Usecase {
	return Usecase{
		AlbumUsecase: albumUsecase,
	}
}
