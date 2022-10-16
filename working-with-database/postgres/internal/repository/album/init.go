package repository

import (
	"database/sql"
	"postgres/internal/entity"
)

type AlbumRepository interface {
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

func NewAlbumRepository(db *sql.DB) AlbumRepository {
	return &albumConnection{db: db}
}
