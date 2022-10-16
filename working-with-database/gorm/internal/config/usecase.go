package config

import albumUsecase "gorm/internal/usecase/album"

type Usecase struct {
	AlbumUsecase albumUsecase.AlbumUsecase
}

func InitUsecase(albumUsecase albumUsecase.AlbumUsecase) Usecase {
	return Usecase{
		AlbumUsecase: albumUsecase,
	}
}
