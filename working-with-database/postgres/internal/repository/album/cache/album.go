package cache

import (
	"context"
	"encoding/json"
	"postgres/internal/entity"

	"github.com/go-redis/redis/v8"
)

// GetAllAlbum is function to get all albums from database
func (repo *albumConnection) GetAllAlbum() ([]entity.Album, error) {
	var albums []entity.Album

	ctx := context.Background()

	albumsString, err := repo.client.Get(ctx, albumsKey).Result()
	if err == redis.Nil {
		return albums, nil
	}
	if err != nil {
		return albums, err
	}

	err = json.Unmarshal([]byte(albumsString), &albums)
	if err != nil {
		return albums, err
	}

	return albums, nil
}

func (repo *albumConnection) SetAllAlbum(albums []entity.Album) error {
	ctx := context.Background()

	albumsString, err := json.Marshal(albums)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, albumsKey, albumsString, expiration).Err(); err != nil {
		return err
	}

	return nil
}
