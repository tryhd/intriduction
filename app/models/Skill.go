package models

import (
	"time"
)

type Skill struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement:true;not null"`
	ProfileID int       `json:"profile_id" gorm:"not null"`
	Skill     string    `json:"skill" gorm:"not null"`
	Level     string    `json:"level" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
