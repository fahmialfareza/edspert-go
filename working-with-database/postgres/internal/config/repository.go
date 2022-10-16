package config

import (
	"database/sql"
	albumRepository "postgres/internal/repository/album"
)

type Repository struct {
	AlbumRepository albumRepository.AlbumRepository
}

func InitRepository(db *sql.DB) Repository {
	return Repository{
		AlbumRepository: albumRepository.NewAlbumRepository(db),
	}
}
