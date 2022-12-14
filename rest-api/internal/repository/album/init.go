package repository

import (
	"database/sql"
	"postgres/internal/entity"
	"postgres/internal/repository/album/cache"
	"postgres/internal/repository/album/psql"

	"github.com/go-redis/redis/v8"
)

type AlbumRepository interface {
	Get(id int64) (*entity.Album, error)
	Create(album *entity.Album) (int64, error)
	GetAllAlbum() ([]entity.Album, error)
	BatchCreate(albums []entity.Album) ([]int64, error)
	Update(album entity.Album) error
	Delete(id int64) error

	GetAlbumCache(id int64) (*entity.Album, error)
	GetAllAlbumCache() ([]entity.Album, error)
	SetAlbumCache(id int64, album entity.Album) error
	SetAllAlbumCache(albums []entity.Album) error
	DeleteAlbumCache(id int64) error
}

type albumRepository struct {
	postgres psql.AlbumPostgres
	cache    cache.AlbumPostgres
}

// The function is to initialize the album repository
func NewAlbumRepository(db *sql.DB, client *redis.Client) AlbumRepository {
	return &albumRepository{
		postgres: psql.NewAlbumPostgres(db),
		cache:    cache.NewAlbumPostgres(client),
	}
}
