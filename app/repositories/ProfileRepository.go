package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	InsertProfile(Profile models.Profile) models.Profile

	FindProfileByID(ProfileID int) models.Profile

	UpdateProfile(Profile models.Profile) models.Profile
}

type profileConnection struct{ connection *gorm.DB }

func NewProfileRepository(dbConn *gorm.DB) ProfileRepository {
	return &profileConnection{connection: dbConn}
}

func (db *profileConnection) InsertProfile(profile models.Profile) models.Profile {

	db.connection.Create(&profile)
	db.connection.Find(&profile)
	return profile
}

func (db *profileConnection) FindProfileByID(profileID int) models.Profile {
	var profile models.Profile
	db.connection.Where("profile_code =?", profileID).First(&profile)
	return profile
}

func (db *profileConnection) UpdateProfile(profile models.Profile) models.Profile {
	db.connection.Omit("created_at").Save(&profile)
	db.connection.Find(&profile)
	return profile
}
