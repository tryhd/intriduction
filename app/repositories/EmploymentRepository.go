package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type EmploymentRepository interface {
	InsertEmployment(employment models.Employment) models.Employment

	FindEmploymentByID(employmentID int) models.Employment

	UpdateEmployment(employment models.Employment) models.Employment
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

func (db *employmentConnection) FindEmploymentByID(employmentID int) models.Employment {
	var employment models.Employment
	db.connection.Where("profile_id =?", employmentID).First(&employment)
	return employment
}

func (db *employmentConnection) UpdateEmployment(employment models.Employment) models.Employment {
	db.connection.Omit("created_at").Save(&employment)
	db.connection.Find(&employment)
	return employment
}
