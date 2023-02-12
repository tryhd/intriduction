package services

import (
	"intoduction/app/dtos"
	"intoduction/app/models"
	"intoduction/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type ProfileService interface {
	Insert(Profile dtos.ProfileCreateDTO) models.Profile

	Update(Profile dtos.ProfileUpdateDTO) models.Profile

	FindByID(ProfileID int) models.Profile
}

type profileService struct {
	profileRepository repositories.ProfileRepository
}

func NewProfileService(profileRepo repositories.ProfileRepository) ProfileService {
	return &profileService{
		profileRepository: profileRepo,
	}
}

func (service *profileService) Insert(Profile dtos.ProfileCreateDTO) models.Profile {
	newProfile := models.Profile{}
	err := smapping.FillStruct(&newProfile, smapping.MapFields(&Profile))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.profileRepository.InsertProfile(newProfile)
	return res
}

func (service *profileService) Update(Profile dtos.ProfileUpdateDTO) models.Profile {
	newProfile := models.Profile{}
	err := smapping.FillStruct(&newProfile, smapping.MapFields(&Profile))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.profileRepository.UpdateProfile(newProfile)
	return res
}

func (service *profileService) FindByID(ProfileID int) models.Profile {
	return service.profileRepository.FindProfileByID(ProfileID)
}
