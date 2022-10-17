package album

import (
	"postgres/internal/entity"
	albumUsecase "postgres/internal/usecase/album"
)

type AlbumHandler interface {
	Get(id int64) (*entity.Album, error)
	Create(album *entity.Album) (*entity.Album, error)
	GetAllAlbum() ([]entity.Album, error)
	BatchCreate(albums []entity.Album) ([]entity.Album, error)
	Update(album entity.Album) (entity.Album, error)
	Delete(id int64) error
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
