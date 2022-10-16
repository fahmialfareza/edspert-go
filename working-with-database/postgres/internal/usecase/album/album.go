package usecase

import "postgres/internal/entity"

// It will call the function Get in album repository
func (usecase *albumUsecase) Get(id int64) (*entity.Album, error) {
	return usecase.albumRepository.Get(id)
}

// It will call the function Create in album repository
func (usecase *albumUsecase) Create(album *entity.Album) error {
	return usecase.albumRepository.Create(album)
}

// It will call the function GetAllAlbum in album repository
func (usecase *albumUsecase) GetAllAlbum() ([]entity.Album, error) {
	return usecase.albumRepository.GetAllAlbum()
}

// It will call the function BatchCreate in album repository
func (usecase *albumUsecase) BatchCreate(albums []entity.Album) error {
	return usecase.albumRepository.BatchCreate(albums)
}

// It will call the function Update in album repository
func (usecase *albumUsecase) Update(album entity.Album) error {
	return usecase.albumRepository.Update(album)
}

// It will call the function Delete in album repository
func (usecase *albumUsecase) Delete(id int64) error {
	return usecase.albumRepository.Delete(id)
}
