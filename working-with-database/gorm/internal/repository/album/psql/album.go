package psql

import (
	"errors"
	"gorm/internal/entity"
	"log"
)

// BulkInsertAlbum is function to insert some albums in once to database
func (repository *albumRepository) BulkInsertAlbum(albums []entity.Album) error {
	// Insert some albums to database
	err := repository.db.CreateInBatches(albums, len(albums))
	if err.Error != nil {
		log.Printf("error when bulk insert err: %v", err)
		return err.Error
	}
	log.Print("Success bulk insert Album")

	return nil
}

// GetAllAlbum is function to get all albums from database
func (repository *albumRepository) GetAllAlbum() ([]entity.Album, error) {
	var album []entity.Album

	// Find some albums and set the result to album variable
	err := repository.db.Find(&album).Error
	if err != nil {
		log.Printf("error when getting album err: %v", err)
		return nil, err
	}

	return album, nil
}

// GetAlbumById is function to get specific album by id from database
func (repository *albumRepository) GetAlbumById(id int64) (entity.Album, error) {
	var album entity.Album

	// Find the specific album and set the result to album variable
	err := repository.db.Where("id = ?", id).Take(&album).Error
	if err != nil {
		log.Printf("error when getting album err: %v", err)
		return entity.Album{}, err
	}

	return album, nil
}

// CreateAlbum is function to create album to database
func (repository *albumRepository) CreateAlbum(album entity.Album) error {
	// 
	tx := repository.db.Create(album)
	if tx.Error != nil {
		return tx.Error
	}
	log.Print("Success create Album")
	return nil
}

// UpdateAlbum is function to update album in database
func (repository *albumRepository) UpdateAlbum(album entity.Album) error {
	var albumUpdate entity.Album

	// Get the album by id from database and set the result to albumUpdate variable
	repository.db.Where("id = ?", album.ID).Find(&albumUpdate)
	if albumUpdate.ID == 0 {
		return errors.New("not found")
	}

	// Set the updated data
	albumUpdate.Title = album.Title
	albumUpdate.Artist = album.Artist
	albumUpdate.Price = album.Price

	// Save changes to database
	err := repository.db.Save(albumUpdate).Error
	if err != nil {
		return err
	}
	log.Print("Success update Album")

	return nil
}

// DeleteAlbum is function to delete album in database
func (repository *albumRepository) DeleteAlbum(id int) error {
	var albumUpdate entity.Album

	// Get the album by id from database and set the result to albumUpdate variable
	repository.db.Where("id = ?", id).Find(&albumUpdate)

	// Delete the album that has been gotten from albumUpdate
	repository.db.Delete(albumUpdate)
	log.Print("Success delete Album")
	
	return nil
}
