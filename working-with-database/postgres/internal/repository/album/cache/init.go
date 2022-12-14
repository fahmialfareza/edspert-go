package cache

import (
	"postgres/internal/entity"

	"github.com/go-redis/redis/v8"
)

type AlbumPostgres interface {
	GetAllAlbum() ([]entity.Album, error)
	SetAllAlbum(albums []entity.Album) error
}

type albumConnection struct {
	client *redis.Client
}

// The function is to initialize the album psql repository
func NewAlbumPostgres(cache *redis.Client) AlbumPostgres {
	return &albumConnection{client: cache}
}
