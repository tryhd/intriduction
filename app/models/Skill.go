package models

import (
	"time"
)

type Skill struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;not null"`
	ProfileID int       `json:"profile_id,omitempty" gorm:"not null"`
	Skill     string    `json:"skill,omitempty" gorm:"not null"`
	Level     string    `json:"level,omitempty" gorm:"not null"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
