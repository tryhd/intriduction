package services

import (
	"intoduction/app/dtos"
	"intoduction/app/models"
	"intoduction/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type EducationService interface {
	Get(id int) []models.Education
	Delete(education models.Education) models.Education
	Insert(education dtos.EducationCreateDTO) models.Education
}

type educationService struct {
	educationRepository repositories.EducationRepository
}

func NewEducationService(educationRepo repositories.EducationRepository) EducationService {
	return &educationService{
		educationRepository: educationRepo,
	}
}

func (service *educationService) Insert(education dtos.EducationCreateDTO) models.Education {
	newEducation := models.Education{}
	err := smapping.FillStruct(&newEducation, smapping.MapFields(&education))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.educationRepository.InsertEducation(newEducation)
	return res
}

func (service *educationService) Delete(education models.Education) models.Education {
	return service.educationRepository.DeleteEducation(education)
}

func (service *educationService) Get(id int) []models.Education {
	return service.educationRepository.GetEducation(id)
}
