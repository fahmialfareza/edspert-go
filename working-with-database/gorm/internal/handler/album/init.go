package album

import (
	"gorm/internal/entity"
	albumUsecase "gorm/internal/usecase/album"
)

type AlbumHandler interface {
	GetAllAlbum() ([]entity.Album, error)
	GetAlbumById(id int64) (entity.Album, error)
	CreateAlbum(album entity.Album) error
	BulkInsertAlbum(albums []entity.Album) error
	UpdateAlbum(album entity.Album) error
	DeleteAlbum(id int) error
}

type albumHandler struct {
	albumUsecase albumUsecase.AlbumUsecase
}

// The function is to initialize the album handler
func NewAlbumHandler(albumUsecase albumUsecase.AlbumUsecase) AlbumHandler {
	return &albumHandler{
		albumUsecase: albumUsecase,
	}
}
