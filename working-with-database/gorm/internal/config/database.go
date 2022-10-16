package config

import (
	"gorm/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// we will define models here
	db.AutoMigrate(&entity.Album{})
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}

	dbSQL.Close()
}
