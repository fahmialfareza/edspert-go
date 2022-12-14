package cache

import (
	"postgres/internal/entity"

	"github.com/go-redis/redis/v8"
)

type AlbumPostgres interface {
	GetAlbum(id int64) (*entity.Album, error)
	GetAllAlbum() ([]entity.Album, error)
	SetAlbum(id int64, album entity.Album) error
	SetAllAlbum(albums []entity.Album) error
	Delete(id int64) error
}

type albumConnection struct {
	client *redis.Client
}

// The function is to initialize the album psql repository
func NewAlbumPostgres(cache *redis.Client) AlbumPostgres {
	return &albumConnection{client: cache}
}
