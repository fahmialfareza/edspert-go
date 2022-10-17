package usecase

import "postgres/internal/entity"

// It will call the function Get in album repository
func (usecase *albumUsecase) Get(id int64) (*entity.Album, error) {
	return usecase.albumRepository.Get(id)
}

// It will call the function Create in album repository
func (usecase *albumUsecase) Create(album *entity.Album) (*entity.Album, error) {
	var newAlbum *entity.Album

	// Create album
	id, err := usecase.albumRepository.Create(album)
	if err != nil {
		return newAlbum, err
	}

	// Find new album
	newAlbum, err = usecase.albumRepository.Get(id)
	if err != nil {
		return newAlbum, err
	}

	return newAlbum, nil
}

// It will call the function GetAllAlbum in album repository
func (usecase *albumUsecase) GetAllAlbum() ([]entity.Album, error) {
	return usecase.albumRepository.GetAllAlbum()
}

// It will call the function BatchCreate in album repository
func (usecase *albumUsecase) BatchCreate(albums []entity.Album) ([]entity.Album, error) {
	var newAlbums []entity.Album

	// Batch create and get the new id
	ids, err := usecase.albumRepository.BatchCreate(albums)
	if err != nil {
		return newAlbums, err
	}

	// Get detail album by ids
	for _, id := range ids {
		album, err := usecase.albumRepository.Get(id)
		if err != nil {
			return newAlbums, err
		}

		newAlbums = append(newAlbums, *album)
	}

	return newAlbums, nil
}

// It will call the function Update in album repository
func (usecase *albumUsecase) Update(album entity.Album) (entity.Album, error) {
	var updatedAlbum *entity.Album

	// Update album
	err := usecase.albumRepository.Update(album)
	if err != nil {
		return *updatedAlbum, err
	}

	// Find new album
	updatedAlbum, err = usecase.albumRepository.Get(album.ID)
	if err != nil {
		return *updatedAlbum, err
	}

	return *updatedAlbum, nil
}

// It will call the function Delete in album repository
func (usecase *albumUsecase) Delete(id int64) error {
	return usecase.albumRepository.Delete(id)
}
