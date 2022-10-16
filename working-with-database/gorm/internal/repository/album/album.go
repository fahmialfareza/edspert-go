package repository

import (
	"gorm/internal/entity"
)

// It will call the function BulkInsertAlbum in psql/album
func (repository *albumRepository) BulkInsertAlbum(albums []entity.Album) error {
	return repository.postgres.BulkInsertAlbum(albums)
}

// It will call the function GetAllAlbum in psql/album
func (repository *albumRepository) GetAllAlbum() ([]entity.Album, error) {
	return repository.postgres.GetAllAlbum()
}

// It will call the function GetAlbumById in psql/album
func (repository *albumRepository) GetAlbumById(id int64) (entity.Album, error) {
	return repository.postgres.GetAlbumById(id)
}

// It will call the function CreateAlbum in psql/album
func (repository *albumRepository) CreateAlbum(album entity.Album) error {
	return repository.postgres.CreateAlbum(album)
}

// It will call the function UpdateAlbum in psql/album
func (repository *albumRepository) UpdateAlbum(album entity.Album) error {
	return repository.postgres.UpdateAlbum(album)
}

// It will call the function DeleteAlbum in psql/album
func (repository *albumRepository) DeleteAlbum(id int) error {
	return repository.postgres.DeleteAlbum(id)
}
