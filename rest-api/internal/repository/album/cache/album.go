package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"postgres/internal/entity"

	"github.com/go-redis/redis/v8"
)

// Get Specific album cache
func (repo *albumConnection) GetAlbum(id int64) (*entity.Album, error) {
	var album entity.Album

	ctx := context.Background()
	key := fmt.Sprintf(albumDetailKey, id)

	albumsString, err := repo.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return &album, nil
	}
	if err != nil {
		return &album, err
	}

	err = json.Unmarshal([]byte(albumsString), &album)
	if err != nil {
		return &album, err
	}

	return &album, nil
}

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

func (repo *albumConnection) SetAlbum(id int64, album entity.Album) error {
	ctx := context.Background()
	key := fmt.Sprintf(albumDetailKey, id)

	albumsString, err := json.Marshal(album)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, key, albumsString, expiration).Err(); err != nil {
		return err
	}

	return nil
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

func (repo *albumConnection) Delete(id int64) error {
	ctx := context.Background()
	key := fmt.Sprintf(albumDetailKey, id)

	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
