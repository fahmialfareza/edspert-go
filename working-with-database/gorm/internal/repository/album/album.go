package repository

import (
	"gorm/internal/entity"
)

func (repository *albumRepository) BulkInsertAlbum(albums []entity.Album) error {
	return repository.postgres.BulkInsertAlbum(albums)
}

func (repository *albumRepository) GetAllAlbum() ([]entity.Album, error) {
	return repository.postgres.GetAllAlbum()
}

func (repository *albumRepository) GetAlbumById(id int64) (entity.Album, error) {
	return repository.postgres.GetAlbumById(id)
}

func (repository *albumRepository) CreateAlbum(album entity.Album) error {
	return repository.postgres.CreateAlbum(album)
}

func (repository *albumRepository) UpdateAlbum(album entity.Album) error {
	return repository.postgres.UpdateAlbum(album)
}

func (repository *albumRepository) DeleteAlbum(id int) error {
	return repository.postgres.DeleteAlbum(id)
}
