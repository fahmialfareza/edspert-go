package album

import (
	"gorm/internal/entity"
)

func (handler *albumHandler) GetAllAlbum() ([]entity.Album, error) {
	return handler.albumUsecase.GetAllAlbum()
}

func (handler *albumHandler) GetAlbumById(id int64) (entity.Album, error) {
	return handler.albumUsecase.GetAlbumById(id)
}

func (handler *albumHandler) CreateAlbum(album entity.Album) error {
	return handler.albumUsecase.CreateAlbum(album)
}

func (handler *albumHandler) BulkInsertAlbum(albums []entity.Album) error {
	return handler.albumUsecase.BulkInsertAlbum(albums)
}

func (handler *albumHandler) UpdateAlbum(album entity.Album) error {
	return handler.albumUsecase.UpdateAlbum(album)
}

func (handler *albumHandler) DeleteAlbum(id int) error {
	return handler.albumUsecase.DeleteAlbum(id)
}
