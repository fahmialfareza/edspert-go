package album

import "postgres/internal/entity"

func (handler *albumHandler) Get(id int64) (*entity.Album, error) {
	return handler.albumUsecase.Get(id)
}

func (handler *albumHandler) Create(album *entity.Album) error {
	return handler.albumUsecase.Create(album)
}

func (handler *albumHandler) GetAllAlbum() ([]entity.Album, error) {
	return handler.albumUsecase.GetAllAlbum()
}

func (handler *albumHandler) BatchCreate(albums []entity.Album) error {
	return handler.albumUsecase.BatchCreate(albums)
}

func (handler *albumHandler) Update(album entity.Album) error {
	return handler.albumUsecase.Update(album)
}

func (handler *albumHandler) Delete(id int64) error {
	return handler.albumUsecase.Delete(id)
}
