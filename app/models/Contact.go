package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Contact struct {
	ID             string    `json:"id" gorm:"uniqueIndex;not null" `
	ContactName    string    `json:"contact_name" gorm:"not null"`
	ContactEmail   string    `json:"contact_email" gorm:"not null"`
	ContactMessage string    `json:"contact_message" gorm:"not null"`
	Status         uint8     `json:"status" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime" `
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime" `
}

func (product *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
