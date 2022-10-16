package psql

import (
	"gorm/internal/entity"

	"gorm.io/gorm"
)

type AlbumRepository interface {
	GetAllAlbum() ([]entity.Album, error)
	GetAlbumById(id int64) (entity.Album, error)
	CreateAlbum(album entity.Album) error
	BulkInsertAlbum(albums []entity.Album) error
	UpdateAlbum(album entity.Album) error
	DeleteAlbum(id int) error
}

type albumRepository struct {
	db *gorm.DB
}

// The function is to initialize the album psql repository
func NewAlbumRepository(database *gorm.DB) AlbumRepository {
	return &albumRepository{
		db: database,
	}
}
