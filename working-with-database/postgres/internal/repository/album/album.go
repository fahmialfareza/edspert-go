package repository

import (
	"postgres/internal/entity"
)

// It will call the function Create in psql/album
func (repo *albumRepository) Create(album *entity.Album) (int64, error) {
	return repo.postgres.Create(album)
}

// It will call the function Get in psql/album
func (repo *albumRepository) Get(id int64) (*entity.Album, error) {
	return repo.postgres.Get(id)
}

// It will call the function GetAllAlbum in psql/album
func (repo *albumRepository) GetAllAlbum() ([]entity.Album, error) {
	return repo.postgres.GetAllAlbum()
}

// It will call the function BatchCreate in psql/album
func (repo *albumRepository) BatchCreate(albums []entity.Album) ([]int64, error) {
	return repo.postgres.BatchCreate(albums)
}

// It will call the function Update in psql/album
func (repo *albumRepository) Update(album entity.Album) error {
	return repo.postgres.Update(album)
}

// It will call the function Delete in psql/album
func (repo *albumRepository) Delete(id int64) error {
	return repo.postgres.Delete(id)
}
