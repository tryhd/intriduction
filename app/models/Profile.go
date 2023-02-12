package models

import (
	"time"
)

type Profile struct {
	ProfileCode    int       `json:"profileCode" gorm:"primaryKey;autoIncrement:true;not null"`
	WantedJobTitle string    `json:"wantedJobTitle"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Country        string    `json:"country"`
	City           string    `json:"city"`
	Address        string    `json:"address"`
	PostalCode     int       `json:"postalCode"`
	DrivingLicense string    `json:"drivingLicense"`
	Nationality    string    `json:"nationality"`
	PlaceOfBirth   string    `json:"placeOfBirth"`
	DateOfBirth    string    `json:"dateOfBirth"`
	PhotoUrl       string    `json:"photoUrl"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
