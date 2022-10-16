package repository

import (
	"postgres/internal/entity"
)

func (repo *albumRepository) Create(album *entity.Album) error {
	return repo.postgres.Create(album)
}

func (repo *albumRepository) Get(id int64) (*entity.Album, error) {
	return repo.postgres.Get(id)
}

func (repo *albumRepository) GetAllAlbum() ([]entity.Album, error) {
	return repo.postgres.GetAllAlbum()
}

func (repo *albumRepository) BatchCreate(albums []entity.Album) error {
	return repo.postgres.BatchCreate(albums)
}

func (repo *albumRepository) Update(album entity.Album) error {
	return repo.postgres.Update(album)
}

func (repo *albumRepository) Delete(id int64) error {
	return repo.postgres.Delete(id)
}
