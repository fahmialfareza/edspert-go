package album

import (
	"gorm/internal/entity"
	albumRepository "gorm/internal/repository/album"
)

type AlbumUsecase interface {
	GetAllAlbum() ([]entity.Album, error)
	GetAlbumById(id int64) (entity.Album, error)
	CreateAlbum(album entity.Album) error
	BulkInsertAlbum(albums []entity.Album) error
	UpdateAlbum(album entity.Album) error
	DeleteAlbum(id int) error
}

type albumUsecase struct {
	albumRepository albumRepository.AlbumRepository
}

func NewAlbumUsecase(albumRepository albumRepository.AlbumRepository) AlbumUsecase {
	return &albumUsecase{
		albumRepository: albumRepository,
	}
}
