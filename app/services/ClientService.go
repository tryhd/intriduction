package services

import (
	"intoduction/app/dtos"
	"intoduction/app/models"
	"intoduction/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type ClientService interface {
	Insert(client dtos.ClientCreateDTO) models.Client
	Update(client dtos.ClientUpdateDTO) models.Client
	Restore(client dtos.ClientRestoreDTO) models.Client
	Delete(client models.Client) models.Client
	DeletePermanent(client models.Client)
	All() []models.Client
	Deleted() []models.Client
	FindByID(clientID string) models.Client
}

type clientService struct {
	clientRepository repositories.ClientRepository
}

func NewClientService(clientRepo repositories.ClientRepository) ClientService {
	return &clientService{
		clientRepository: clientRepo,
	}
}

func (service *clientService) Insert(client dtos.ClientCreateDTO) models.Client {
	newClient := models.Client{}
	err := smapping.FillStruct(&newClient, smapping.MapFields(&client))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.clientRepository.InsertClient(newClient)
	return res
}

func (service *clientService) Update(client dtos.ClientUpdateDTO) models.Client {
	newClient := models.Client{}
	err := smapping.FillStruct(&newClient, smapping.MapFields(&client))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.clientRepository.UpdateClient(newClient)
	return res
}

func (service *clientService) Delete(client models.Client) models.Client {
	return service.clientRepository.DeleteClient(client)
}

func (service *clientService) All() []models.Client {
	return service.clientRepository.AllClient()
}

func (service *clientService) FindByID(clientID string) models.Client {
	return service.clientRepository.FindClientByID(clientID)
}

func (service *clientService) Deleted() []models.Client {
	return service.clientRepository.DeletedClient()
}

func (service *clientService) Restore(client dtos.ClientRestoreDTO) models.Client {
	newClient := models.Client{}
	err := smapping.FillStruct(&newClient, smapping.MapFields(&client))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.clientRepository.RestoreClient(newClient)
	return res
}

func (service *clientService) DeletePermanent(client models.Client) {
	service.clientRepository.DeletePermanentClient(client)
}
