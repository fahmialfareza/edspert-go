package usecase

import "postgres/internal/entity"

// It will call the function Get in album repository
func (usecase *albumUsecase) Get(id int64) (*entity.Album, error) {
	// Get from cache
	album, err := usecase.albumRepository.GetAlbumCache(id)
	if err != nil {
		return album, err
	}

	if album.ID != 0 {
		return album, nil
	}

	// Get from db
	album, err = usecase.albumRepository.Get(id)
	if err != nil {
		return album, err
	}

	// Set to cache
	if err = usecase.albumRepository.SetAlbumCache(id, *album); err != nil {
		return album, err
	}

	return album, nil
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

	// Find all albums
	albums, err := usecase.albumRepository.GetAllAlbum()
	if err != nil {
		return newAlbum, err
	}

	// Set to specific cache
	if err = usecase.albumRepository.SetAlbumCache(id, *newAlbum); err != nil {
		return newAlbum, err
	}

	// Set all cache
	if err = usecase.albumRepository.SetAllAlbumCache(albums); err != nil {
		return newAlbum, err
	}

	return newAlbum, nil
}

// It will call the function GetAllAlbum in album repository
func (usecase *albumUsecase) GetAllAlbum() ([]entity.Album, error) {
	var albums []entity.Album

	// Get from cache
	albums, err := usecase.albumRepository.GetAllAlbumCache()
	if err != nil {
		return albums, err
	}

	if len(albums) > 0 {
		return albums, nil
	}

	// Get from db
	albums, err = usecase.albumRepository.GetAllAlbum()
	if err != nil {
		return albums, err
	}

	// Set to cache
	if err = usecase.albumRepository.SetAllAlbumCache(albums); err != nil {
		return albums, err
	}

	return albums, nil
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
		// Get from db
		album, err := usecase.albumRepository.Get(id)
		if err != nil {
			return newAlbums, err
		}

		// Set to specific cache
		if err = usecase.albumRepository.SetAlbumCache(id, *album); err != nil {
			return newAlbums, err
		}

		newAlbums = append(newAlbums, *album)
	}

	// Find all albums
	allAlbums, err := usecase.albumRepository.GetAllAlbum()
	if err != nil {
		return newAlbums, err
	}

	// Set all cache
	if err = usecase.albumRepository.SetAllAlbumCache(allAlbums); err != nil {
		return newAlbums, err
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

	// Find all albums
	albums, err := usecase.albumRepository.GetAllAlbum()
	if err != nil {
		return *updatedAlbum, err
	}

	// Set to specific cache
	if err = usecase.albumRepository.SetAlbumCache(album.ID, *updatedAlbum); err != nil {
		return *updatedAlbum, err
	}

	// Set all cache
	if err = usecase.albumRepository.SetAllAlbumCache(albums); err != nil {
		return *updatedAlbum, err
	}

	return *updatedAlbum, nil
}

// It will call the function Delete in album repository
func (usecase *albumUsecase) Delete(id int64) error {
	// Delete from db
	if err := usecase.albumRepository.Delete(id); err != nil {
		return err
	}

	// Delete from cache
	if err := usecase.albumRepository.DeleteAlbumCache(id); err != nil {
		return err
	}

	// Find all albums
	albums, err := usecase.albumRepository.GetAllAlbum()
	if err != nil {
		return err
	}

	// Set all cache
	if err = usecase.albumRepository.SetAllAlbumCache(albums); err != nil {
		return err
	}

	return nil
}
