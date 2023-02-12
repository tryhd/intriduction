package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ID          string         `json:"id" gorm:"uniqueIndex;not null"`
	ClientName  string         `json:"client_name" gorm:"not null"`
	ClientImage string         `json:"client_image" gorm:"not null"`
	ClientTitle string         `json:"client_title" gorm:"not null"`
	Description string         `json:"description"`
	Status      uint8          `json:"status" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (product *Client) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}
