package usecase

import (
	"postgres/internal/entity"
	albumRepository "postgres/internal/repository/album"
)

type AlbumUsecase interface {
	Get(id int64) (*entity.Album, error)
	Create(album *entity.Album) error
	GetAllAlbum() ([]entity.Album, error)
	BatchCreate(albums []entity.Album) error
	Update(album entity.Album) error
	Delete(id int64) error
}

type albumUsecase struct {
	albumRepository albumRepository.AlbumRepository
}

// The function is to initialize the album usecase
func NewAlbumUsecase(albumRepository albumRepository.AlbumRepository) AlbumUsecase {
	return &albumUsecase{
		albumRepository: albumRepository,
	}
}
