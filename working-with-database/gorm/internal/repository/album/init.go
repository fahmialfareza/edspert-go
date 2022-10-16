package repository

import (
	"gorm/internal/entity"
	"gorm/internal/repository/album/psql"

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
	postgres psql.AlbumRepository
}

// The function is to initialize the album repository
func NewAlbumRepository(db *gorm.DB) AlbumRepository {
	return &albumRepository{
		postgres: psql.NewAlbumRepository(db),
	}
}
