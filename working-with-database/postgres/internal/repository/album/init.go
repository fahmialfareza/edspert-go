package repository

import (
	"database/sql"
	"postgres/internal/entity"
	"postgres/internal/repository/album/psql"
)

type AlbumRepository interface {
	Get(id int64) (*entity.Album, error)
	Create(album *entity.Album) (int64, error)
	GetAllAlbum() ([]entity.Album, error)
	BatchCreate(albums []entity.Album) ([]int64, error)
	Update(album entity.Album) error
	Delete(id int64) error
}

type albumRepository struct {
	postgres psql.AlbumPostgres
}

// The function is to initialize the album repository
func NewAlbumRepository(db *sql.DB) AlbumRepository {
	return &albumRepository{
		postgres: psql.NewAlbumPostgres(db),
	}
}
