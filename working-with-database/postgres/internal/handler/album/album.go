package album

import "postgres/internal/entity"

// It will call the function Get in album usecase
func (handler *albumHandler) Get(id int64) (*entity.Album, error) {
	return handler.albumUsecase.Get(id)
}

// It will call the function Create in album usecase
func (handler *albumHandler) Create(album *entity.Album) (*entity.Album, error) {
	return handler.albumUsecase.Create(album)
}

// It will call the function GetAllAlbum in album usecase
func (handler *albumHandler) GetAllAlbum() ([]entity.Album, error) {
	return handler.albumUsecase.GetAllAlbum()
}

// It will call the function BatchCreate in album usecase
func (handler *albumHandler) BatchCreate(albums []entity.Album) ([]entity.Album, error) {
	return handler.albumUsecase.BatchCreate(albums)
}

// It will call the function Update in album usecase
func (handler *albumHandler) Update(album entity.Album) (entity.Album, error) {
	return handler.albumUsecase.Update(album)
}

// It will call the function Delete in album usecase
func (handler *albumHandler) Delete(id int64) error {
	return handler.albumUsecase.Delete(id)
}
