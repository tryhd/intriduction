package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type EmploymentRepository interface {
	InsertEmployment(employment models.Employment) models.Employment

	GetEmployment(id int) []models.Employment

	DeleteEmployment(employment models.Employment) models.Employment
}

type employmentConnection struct{ connection *gorm.DB }

func NewEmploymentRepository(dbConn *gorm.DB) EmploymentRepository {
	return &employmentConnection{connection: dbConn}
}

func (db *employmentConnection) InsertEmployment(employment models.Employment) models.Employment {

	db.connection.Create(&employment)
	db.connection.Find(&employment)
	return employment
}

func (db *employmentConnection) GetEmployment(id int) []models.Employment {
	var employments []models.Employment
	db.connection.Select("id", "job_title", "employer", "start_date", "end_date", "city", "description").Omit("created_at", "updated_at").Where("profile_id", id).Find(&employments)
	return employments
}

func (db *employmentConnection) DeleteEmployment(employment models.Employment) models.Employment {
	db.connection.Debug().Where("profile_id", employment.ProfileID).Where("id", employment.ID).Delete(&employment)
	db.connection.Select("profile_id").Find(&employment)
	res := models.Employment{
		ID:        employment.ID,
		ProfileID: employment.ProfileID,
	}
	return res
}
