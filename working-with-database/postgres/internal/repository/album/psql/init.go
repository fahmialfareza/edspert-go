package psql

import (
	"database/sql"
	"postgres/internal/entity"
)

type AlbumPostgres interface {
	Get(id int64) (*entity.Album, error)
	Create(album *entity.Album) error
	GetAllAlbum() ([]entity.Album, error)
	BatchCreate(albums []entity.Album) error
	Update(album entity.Album) error
	Delete(id int64) error
}

type albumConnection struct {
	db *sql.DB
}

// The function is to initialize the album psql repository
func NewAlbumPostgres(db *sql.DB) AlbumPostgres {
	return &albumConnection{db: db}
}
