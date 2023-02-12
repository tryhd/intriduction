package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Portfolio struct {
	ID                   string         `json:"id" gorm:"uniqueIndex;not null"`
	PortfolioTitle       string         `json:"portfolio_title" gorm:"not null"`
	PortfolioImage       string         `json:"portfolio_image" gorm:"not null"`
	PortfolioDescription string         `json:"portfolio_description" gorm:"not null"`
	PortfolioLink        string         `json:"portfolio_link"`
	Status               uint8          `json:"status" gorm:"not null"`
	CreatedAt            time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt            time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt            gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (product *Portfolio) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
