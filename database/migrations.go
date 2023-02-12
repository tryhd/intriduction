package database

import (
	"intoduction/app/models"
	"intoduction/configs"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = configs.SetupDatabaseConnection()
)

func InitialMigration() {
	db.Migrator().DropTable(&models.Profile{})
	db.Migrator().CreateTable(&models.Profile{})
}
