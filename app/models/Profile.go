package models

import (
	"time"
)

type Profile struct {
	ProfileCode       int                 `json:"profileCode,omitempty" gorm:"primaryKey;autoIncrement:true;not null"`
	WantedJobTitle    string              `json:"wantedJobTitle,omitempty"`
	FirstName         string              `json:"firstName,omitempty"`
	LastName          string              `json:"lastName,omitempty"`
	Email             string              `json:"email,omitempty"`
	Phone             string              `json:"phone,omitempty"`
	Country           string              `json:"country,omitempty"`
	City              string              `json:"city,omitempty"`
	Address           string              `json:"address,omitempty"`
	PostalCode        int                 `json:"postalCode,omitempty"`
	DrivingLicense    string              `json:"drivingLicense,omitempty"`
	Nationality       string              `json:"nationality,omitempty"`
	PlaceOfBirth      string              `json:"placeOfBirth,omitempty"`
	DateOfBirth       string              `json:"dateOfBirth,omitempty"`
	PhotoUrl          string              `json:"photoUrl,omitempty"`
	Skill             []Skill             `json:"skill,omitempty"`
	WorkingExperience []WorkingExperience `json:"working_experience,omitempty"`
	CreatedAt         time.Time           `json:"created_at,omitempty" gorm:"autoCreateTime,omitempty"`
	UpdatedAt         time.Time           `json:"updated_at,omitempty" gorm:"autoUpdateTime,omitempty"`
}
