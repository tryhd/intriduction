package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type ContentRepository interface {
	InsertContent(cl models.Content) models.Content
	UpdateContent(cl models.Content) models.Content
	DeleteContent(cl models.Content) models.Content
	AllContent() []models.Content
	FindContentByID(ContentID string) models.Content
}

type contentConnection struct{ connection *gorm.DB }

func NewContentRepository(dbConn *gorm.DB) ContentRepository {
	return &contentConnection{connection: dbConn}
}

func (db *contentConnection) InsertContent(content models.Content) models.Content {
	db.connection.Create(&content)
	db.connection.Find(&content)
	return content
}

func (db *contentConnection) UpdateContent(content models.Content) models.Content {
	db.connection.Save(&content)
	db.connection.Find(&content)
	return content
}

func (db *contentConnection) DeleteContent(content models.Content) models.Content {
	db.connection.Delete(&content)
	db.connection.Find(&content)
	return content
}

func (db *contentConnection) AllContent() []models.Content {
	var contents []models.Content
	db.connection.Find(&contents)
	return contents
}

func (db *contentConnection) FindContentByID(contentID string) models.Content {
	var content models.Content
	db.connection.Where("id =?", contentID).First(&content)
	return content
}
