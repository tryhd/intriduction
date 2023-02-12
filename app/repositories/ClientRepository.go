package repositories

import (
	// "github.com/ydhnwb/golang_api/entity"

	"intoduction/app/models"

	"gorm.io/gorm"
)

type ClientRepository interface {
	InsertClient(cl models.Client) models.Client
	UpdateClient(cl models.Client) models.Client
	DeleteClient(cl models.Client) models.Client
	AllClient() []models.Client
	FindClientByID(clientID string) models.Client
	DeletedClient() []models.Client
	RestoreClient(cl models.Client) models.Client
	DeletePermanentClient(cl models.Client)
}

type clientConnection struct {
	connection *gorm.DB
}

func NewClientRepository(dbConn *gorm.DB) ClientRepository {
	return &clientConnection{
		connection: dbConn,
	}
}

func (db *clientConnection) InsertClient(client models.Client) models.Client {
	db.connection.Create(&client)
	db.connection.Find(&client)
	return client
}

func (db *clientConnection) UpdateClient(client models.Client) models.Client {
	db.connection.Save(&client)
	db.connection.Find(&client)
	return client
}

func (db *clientConnection) DeleteClient(client models.Client) models.Client {
	db.connection.Delete(&client)
	db.connection.Find(&client)
	return client
}

func (db *clientConnection) AllClient() []models.Client {
	var clients []models.Client
	db.connection.Find(&clients)
	return clients
}

func (db *clientConnection) FindClientByID(clientID string) models.Client {
	var client models.Client
	db.connection.Where("id =?", clientID).First(&client)
	return client
}

func (db *clientConnection) DeletedClient() []models.Client {
	var clients []models.Client
	db.connection.Unscoped().Where("deleted_at != 0").Find(&clients)
	return clients
}

func (db *clientConnection) RestoreClient(client models.Client) models.Client {
	db.connection.Unscoped().First(&client).Update("deleted_at", nil)
	db.connection.Find(&client)
	return client
}

func (db *clientConnection) DeletePermanentClient(client models.Client) {
	db.connection.Unscoped().Delete(&client)
}
