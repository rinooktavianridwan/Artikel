package database

import (
	"artikel/artikel"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) error {
	error := db.AutoMigrate(&artikel.Artikel{})
	if error != nil {
		log.Fatal("Failed to migrate database:", error)
	}
	return error
}
