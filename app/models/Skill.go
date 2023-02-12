package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Skill struct {
	ID        string    `json:"id" gorm:"uniqueIndex;not null"`
	SkillName string    `json:"skill_name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (product *Skill) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
