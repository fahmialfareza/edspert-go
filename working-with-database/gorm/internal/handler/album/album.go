package album

import (
	"gorm/internal/entity"
)

// It will call the function GetAllAlbum in album usecase
func (handler *albumHandler) GetAllAlbum() ([]entity.Album, error) {
	return handler.albumUsecase.GetAllAlbum()
}

// It will call the function GetAlbumById in album usecase
func (handler *albumHandler) GetAlbumById(id int64) (entity.Album, error) {
	return handler.albumUsecase.GetAlbumById(id)
}

// It will call the function CreateAlbum in album usecase
func (handler *albumHandler) CreateAlbum(album entity.Album) error {
	return handler.albumUsecase.CreateAlbum(album)
}

// It will call the function BulkInsertAlbum in album usecase
func (handler *albumHandler) BulkInsertAlbum(albums []entity.Album) error {
	return handler.albumUsecase.BulkInsertAlbum(albums)
}

// It will call the function UpdateAlbum in album usecase
func (handler *albumHandler) UpdateAlbum(album entity.Album) error {
	return handler.albumUsecase.UpdateAlbum(album)
}

// It will call the function DeleteAlbum in album usecase
func (handler *albumHandler) DeleteAlbum(id int) error {
	return handler.albumUsecase.DeleteAlbum(id)
}
