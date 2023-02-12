package services

import (
	"intoduction/app/dtos"
	"intoduction/app/models"
	"intoduction/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type SkillService interface {
	Insert(skill dtos.SkillCreateDTO) models.Skill
	Delete(skill models.Skill) models.Skill
	All() []models.Skill
	FindByID(skillID string) models.Skill
}

type skillService struct {
	skillRepository repositories.SkillRepository
}

func NewSkillService(skillRepo repositories.SkillRepository) SkillService {
	return &skillService{
		skillRepository: skillRepo,
	}
}

func (service *skillService) Insert(skill dtos.SkillCreateDTO) models.Skill {
	newSkill := models.Skill{}
	err := smapping.FillStruct(&newSkill, smapping.MapFields(&skill))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.skillRepository.InsertSkill(newSkill)
	return res
}

func (service *skillService) Delete(skill models.Skill) models.Skill {
	return service.skillRepository.DeleteSkill(skill)
}

func (service *skillService) All() []models.Skill {
	return service.skillRepository.AllSkill()
}

func (service *skillService) FindByID(skillID string) models.Skill {
	return service.skillRepository.FindSkillByID(skillID)
}
