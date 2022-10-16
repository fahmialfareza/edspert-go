package psql

import (
	"errors"
	"gorm/internal/entity"
	"log"
)

func (repository *albumRepository) BulkInsertAlbum(albums []entity.Album) error {
	err := repository.db.CreateInBatches(albums, len(albums))
	if err.Error != nil {
		log.Printf("error when bulk insert err: %v", err)
		return err.Error
	}

	log.Print("Success bulk insert Album")
	return nil
}

func (repository *albumRepository) GetAllAlbum() ([]entity.Album, error) {
	var album []entity.Album
	err := repository.db.Find(&album).Error
	if err != nil {
		log.Printf("error when getting album err: %v", err)
		return nil, err
	}
	return album, nil
}

func (repository *albumRepository) GetAlbumById(id int64) (entity.Album, error) {
	var album entity.Album
	err := repository.db.Where("id = ?", id).Take(&album).Error
	if err != nil {
		log.Printf("error when getting album err: %v", err)
		return entity.Album{}, err
	}
	return album, nil
}

func (repository *albumRepository) CreateAlbum(album entity.Album) error {
	tx := repository.db.Create(album)
	if tx.Error != nil {
		return tx.Error
	}
	log.Print("Success create Album")
	return nil
}

func (repository *albumRepository) UpdateAlbum(album entity.Album) error {
	var albumUpdate entity.Album
	repository.db.Where("id = ?", album.ID).Find(&albumUpdate)
	if albumUpdate.ID == 0 {
		return errors.New("not found")
	}

	albumUpdate.Title = album.Title
	albumUpdate.Artist = album.Artist
	albumUpdate.Price = album.Price

	err := repository.db.Save(albumUpdate).Error
	if err != nil {
		return err
	}
	log.Print("Success update Album")
	return nil
}

func (repository *albumRepository) DeleteAlbum(id int) error {
	var albumUpdate entity.Album
	repository.db.Where("id = ?", id).Find(&albumUpdate)
	repository.db.Delete(albumUpdate)
	log.Print("Success delete Album")
	return nil
}
