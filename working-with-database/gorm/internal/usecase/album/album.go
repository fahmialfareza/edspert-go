package album

import "gorm/internal/entity"

// It will call the function GetAllAlbum in album repository
func (usecase *albumUsecase) GetAllAlbum() ([]entity.Album, error) {
	return usecase.albumRepository.GetAllAlbum()
}

// It will call the function GetAlbumById in album repository
func (usecase *albumUsecase) GetAlbumById(id int64) (entity.Album, error) {
	return usecase.albumRepository.GetAlbumById(id)
}

// It will call the function CreateAlbum in album repository
func (usecase *albumUsecase) CreateAlbum(album entity.Album) error {
	return usecase.albumRepository.CreateAlbum(album)
}

// It will call the function BulkInsertAlbum in album repository
func (usecase *albumUsecase) BulkInsertAlbum(albums []entity.Album) error {
	return usecase.albumRepository.BulkInsertAlbum(albums)
}

// It will call the function UpdateAlbum in album repository
func (usecase *albumUsecase) UpdateAlbum(album entity.Album) error {
	return usecase.albumRepository.UpdateAlbum(album)
}

// It will call the function DeleteAlbum in album repository
func (usecase *albumUsecase) DeleteAlbum(id int) error {
	return usecase.albumRepository.DeleteAlbum(id)
}
