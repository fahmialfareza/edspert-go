package config

import (
	albumRepository "gorm/internal/repository/album"

	"gorm.io/gorm"
)

type Repository struct {
	AlbumRepository albumRepository.AlbumRepository
}

// Function to initialize repository
func InitRepository(db *gorm.DB) Repository {
	return Repository{
		AlbumRepository: albumRepository.NewAlbumRepository(db),
	}
}
