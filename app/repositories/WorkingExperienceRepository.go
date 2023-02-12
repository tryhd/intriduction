package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type WorkingExperienceRepository interface {
	InsertWorkingExperience(WorkingExperience models.WorkingExperience) models.WorkingExperience

	FindWorkingExperienceByID(WorkingExperienceID int) models.WorkingExperience

	UpdateWorkingExperience(WorkingExperience models.WorkingExperience) models.WorkingExperience
}

type workingExperienceConnection struct{ connection *gorm.DB }

func NewWorkingExperienceRepository(dbConn *gorm.DB) WorkingExperienceRepository {
	return &workingExperienceConnection{connection: dbConn}
}

func (db *workingExperienceConnection) InsertWorkingExperience(workingExperience models.WorkingExperience) models.WorkingExperience {
	db.connection.Create(&workingExperience)
	db.connection.Find(&workingExperience)
	return workingExperience
}

func (db *workingExperienceConnection) FindWorkingExperienceByID(profileCode int) models.WorkingExperience {
	var workingExperience models.WorkingExperience
	db.connection.Debug().Omit("profiles").Select("working_experience").Where("profile_id =?", profileCode).Find(&workingExperience)
	return workingExperience
}

func (db *workingExperienceConnection) UpdateWorkingExperience(workingExperience models.WorkingExperience) models.WorkingExperience {
	db.connection.Omit("created_at").Save(&workingExperience)
	db.connection.Find(&workingExperience)
	return workingExperience
}
