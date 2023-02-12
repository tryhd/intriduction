package models

import (
	"time"
)

type Employment struct {
	ID          int       `json:"id" gorm:"primaryKey;not null;autoIncrement:true"`
	ProfileID   int       `json:"profileId" gorm:"not null"  binding:"required"`
	JobTitle    string    `json:"jobTitle" gorm:"not null"  binding:"required"`
	Employer    string    `json:"employer" gorm:"not null"  binding:"required"`
	StartDate   string    `json:"startDate" gorm:"not null"  binding:"required"`
	EndDate     string    `json:"endDate" gorm:"not null"  binding:"required"`
	City        string    `json:"city" gorm:"not null"  binding:"required"`
	Description string    `json:"description" gorm:"not null"  binding:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
