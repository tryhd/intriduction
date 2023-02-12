package services

import (
	"intoduction/app/dtos"
	"intoduction/app/models"
	"intoduction/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type WorkingExperienceService interface {
	Insert(workingExperience dtos.WorkingExperienceCreateDTO) models.WorkingExperience

	Update(workingExperience dtos.WorkingExperienceUpdateDTO) models.WorkingExperience

	FindByID(workingExperienceID int) models.WorkingExperience
}

type workingExperienceService struct {
	workingExperienceRepository repositories.WorkingExperienceRepository
}

func NewWorkingExperienceService(workingExperienceRepo repositories.WorkingExperienceRepository) WorkingExperienceService {
	return &workingExperienceService{
		workingExperienceRepository: workingExperienceRepo,
	}
}

func (service *workingExperienceService) Insert(workingExperience dtos.WorkingExperienceCreateDTO) models.WorkingExperience {
	newWorkingExperience := models.WorkingExperience{}
	err := smapping.FillStruct(&newWorkingExperience, smapping.MapFields(&workingExperience))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.workingExperienceRepository.InsertWorkingExperience(newWorkingExperience)
	return res
}

func (service *workingExperienceService) Update(workingExperience dtos.WorkingExperienceUpdateDTO) models.WorkingExperience {
	newWorkingExperience := models.WorkingExperience{}
	err := smapping.FillStruct(&newWorkingExperience, smapping.MapFields(&workingExperience))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.workingExperienceRepository.UpdateWorkingExperience(newWorkingExperience)
	return res
}

func (service *workingExperienceService) FindByID(workingExperienceID int) models.WorkingExperience {
	return service.workingExperienceRepository.FindWorkingExperienceByID(workingExperienceID)
}
