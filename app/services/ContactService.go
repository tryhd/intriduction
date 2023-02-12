package services

import (
	"intoduction/app/dtos"
	"intoduction/app/models"
	"intoduction/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type ContactService interface {
	Insert(contact dtos.ContactCreateDTO) models.Contact
	Update(contact dtos.ContactUpdateDTO) models.Contact
	Delete(contact models.Contact) models.Contact
	All() []models.Contact
	FindByID(contactID string) models.Contact
}

type contactService struct {
	contactRepository repositories.ContactRepository
}

func NewContactService(contactRepo repositories.ContactRepository) ContactService {
	return &contactService{
		contactRepository: contactRepo,
	}
}

func (service *contactService) Insert(contact dtos.ContactCreateDTO) models.Contact {
	newContact := models.Contact{}
	err := smapping.FillStruct(&newContact, smapping.MapFields(&contact))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.contactRepository.InsertContact(newContact)
	return res
}

func (service *contactService) Update(contact dtos.ContactUpdateDTO) models.Contact {
	newContact := models.Contact{}
	err := smapping.FillStruct(&newContact, smapping.MapFields(&contact))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.contactRepository.UpdateContact(newContact)
	return res
}

func (service *contactService) Delete(contact models.Contact) models.Contact {
	return service.contactRepository.DeleteContact(contact)
}

func (service *contactService) All() []models.Contact {
	return service.contactRepository.AllContact()
}

func (service *contactService) FindByID(contactID string) models.Contact {
	return service.contactRepository.FindContactByID(contactID)
}
