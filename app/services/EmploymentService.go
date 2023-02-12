package services

import (
	"intoduction/app/dtos"
	"intoduction/app/models"
	"intoduction/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type EmploymentService interface {
	FindByID(EmploymentID int) models.Employment
	Insert(employment dtos.EmploymentCreateDTO) models.Employment
	Update(employment dtos.EmploymentUpdateDTO) models.Employment
}

type employmentService struct {
	employmentRepository repositories.EmploymentRepository
}

func NewEmploymentService(employmentRepo repositories.EmploymentRepository) EmploymentService {
	return &employmentService{
		employmentRepository: employmentRepo,
	}
}

func (service *employmentService) Insert(employment dtos.EmploymentCreateDTO) models.Employment {
	newEmployment := models.Employment{}
	err := smapping.FillStruct(&newEmployment, smapping.MapFields(&employment))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.employmentRepository.InsertEmployment(newEmployment)
	return res
}

func (service *employmentService) Update(employment dtos.EmploymentUpdateDTO) models.Employment {
	newEmployment := models.Employment{}
	err := smapping.FillStruct(&newEmployment, smapping.MapFields(&employment))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.employmentRepository.UpdateEmployment(newEmployment)
	return res
}

func (service *employmentService) FindByID(EmploymentID int) models.Employment {
	return service.employmentRepository.FindEmploymentByID(EmploymentID)
}
