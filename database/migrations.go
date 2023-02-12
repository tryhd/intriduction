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
	db.Migrator().DropTable(&models.WorkingExperience{})
	db.Migrator().DropTable(&models.Profile{})
	db.Migrator().DropTable(&models.Employment{})

	db.Migrator().CreateTable(&models.Profile{})
	db.Migrator().CreateTable(&models.WorkingExperience{})
	db.Migrator().CreateTable(&models.Employment{})

}
