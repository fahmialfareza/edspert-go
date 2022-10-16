package config

import (
	"gorm/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// The function to connect database (Postgres)
func OpenDB(dsn string) *gorm.DB {
	// Connect the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto migration that means, entity will be automatically be a table in database
	db.AutoMigrate(&entity.Album{})

	return db
}

// The function to close database connection
func CloseDB(db *gorm.DB) {
	// Close the database connection
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}

	dbSQL.Close()
}
