package config

import albumUsecase "postgres/internal/usecase/album"

type Usecase struct {
	AlbumUsecase albumUsecase.AlbumUsecase
}

func InitUsecase(albumUsecase albumUsecase.AlbumUsecase) Usecase {
	return Usecase{
		AlbumUsecase: albumUsecase,
	}
}
