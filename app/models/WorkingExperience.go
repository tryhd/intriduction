package models

import (
	"time"
)

type WorkingExperience struct {
	ID                int       `json:"id" gorm:"primaryKey;not null;autoIncrement:true"`
	ProfileID         int       `json:"profileId" gorm:"not null"  binding:"required"`
	WorkingExperience string    `json:"workingExperience" gorm:"not null"  binding:"required"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type WorkingExperienceGet struct {
	WorkingExperience string `json:"workingExperience" gorm:"not null"  binding:"required"`
}
