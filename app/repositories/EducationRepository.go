package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type EducationRepository interface {
	DeleteEducation(education models.Education) models.Education
	GetEducation(id int) []models.Education
	InsertEducation(education models.Education) models.Education
}

type educationConnection struct{ connection *gorm.DB }

func NewEducationRepository(dbConn *gorm.DB) EducationRepository {
	return &educationConnection{connection: dbConn}
}

func (db *educationConnection) InsertEducation(education models.Education) models.Education {

	db.connection.Create(&education)
	db.connection.Find(&education)
	return education
}

func (db *educationConnection) GetEducation(id int) []models.Education {
	var educations []models.Education
	db.connection.Select("id", "school", "degree", "start_date", "end_date", "city", "description").Omit("created_at", "updated_at").Where("profile_id", id).Find(&educations)
	return educations
}

func (db *educationConnection) DeleteEducation(education models.Education) models.Education {
	db.connection.Debug().Where("profile_id", education.ProfileID).Where("id", education.ID).Delete(&education)
	db.connection.Select("profile_id").Find(&education)
	res := models.Education{
		ID:        education.ID,
		ProfileID: education.ProfileID,
	}
	return res
}
