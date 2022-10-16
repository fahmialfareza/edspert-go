package album

import "gorm/internal/entity"

func (usecase *albumUsecase) GetAllAlbum() ([]entity.Album, error) {
	return usecase.albumRepository.GetAllAlbum()
}

func (usecase *albumUsecase) GetAlbumById(id int64) (entity.Album, error) {
	return usecase.albumRepository.GetAlbumById(id)
}

func (usecase *albumUsecase) CreateAlbum(album entity.Album) error {
	return usecase.albumRepository.CreateAlbum(album)
}

func (usecase *albumUsecase) BulkInsertAlbum(albums []entity.Album) error {
	return usecase.albumRepository.BulkInsertAlbum(albums)
}

func (usecase *albumUsecase) UpdateAlbum(album entity.Album) error {
	return usecase.albumRepository.UpdateAlbum(album)
}

func (usecase *albumUsecase) DeleteAlbum(id int) error {
	return usecase.albumRepository.DeleteAlbum(id)
}
