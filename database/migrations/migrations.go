package migrations

import (
	"github.com/AlanaPeres/biblioteca/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}