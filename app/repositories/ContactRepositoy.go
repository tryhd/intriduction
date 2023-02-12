package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type ContactRepository interface {
	InsertContact(ct models.Contact) models.Contact
	UpdateContact(ct models.Contact) models.Contact
	DeleteContact(ct models.Contact) models.Contact
	AllContact() []models.Contact
	FindContactByID(contactId string) models.Contact
}

type contactConnection struct {
	connection *gorm.DB
}

func NewContactRepository(dbConn *gorm.DB) ContactRepository {
	return &contactConnection{
		connection: dbConn,
	}
}

func (db *contactConnection) InsertContact(contact models.Contact) models.Contact {
	db.connection.Create(&contact)
	db.connection.Find(&contact)
	return contact
}

func (db *contactConnection) UpdateContact(contact models.Contact) models.Contact {
	db.connection.Save(&contact)
	db.connection.Find(&contact)
	return contact
}

func (db *contactConnection) DeleteContact(contact models.Contact) models.Contact {
	db.connection.Delete(&contact)
	db.connection.Find(&contact)
	return contact
}

func (db *contactConnection) AllContact() []models.Contact {
	var contacts []models.Contact
	db.connection.Find(&contacts)
	return contacts
}

func (db *contactConnection) FindContactByID(contactID string) models.Contact {
	var contact models.Contact
	db.connection.Where("id =?", contactID).First(&contact)
	return contact
}
