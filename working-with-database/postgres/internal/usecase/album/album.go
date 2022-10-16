package usecase

import "postgres/internal/entity"

func (usecase *albumUsecase) Get(id int64) (*entity.Album, error) {
	return usecase.albumRepository.Get(id)
}

func (usecase *albumUsecase) Create(album *entity.Album) error {
	return usecase.albumRepository.Create(album)
}

func (usecase *albumUsecase) GetAllAlbum() ([]entity.Album, error) {
	return usecase.albumRepository.GetAllAlbum()
}

func (usecase *albumUsecase) BatchCreate(albums []entity.Album) error {
	return usecase.albumRepository.BatchCreate(albums)
}

func (usecase *albumUsecase) Update(album entity.Album) error {
	return usecase.albumRepository.Update(album)
}

func (usecase *albumUsecase) Delete(id int64) error {
	return usecase.albumRepository.Delete(id)
}
