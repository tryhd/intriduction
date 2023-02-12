package services

import (
	"intoduction/app/helpers"
	"intoduction/app/models"
	"intoduction/app/repositories"
)

type PhotoService interface {
	Insert(photo string, id int) models.Profile
	Get(id int) string
	Delete(id int) models.Profile
}

type photoService struct {
	photoRepository repositories.PhotoRepository
}

func NewPhotoService(photoRepo repositories.PhotoRepository) PhotoService {
	return &photoService{
		photoRepository: photoRepo,
	}
}

func (service *photoService) Insert(photo string, id int) models.Profile {
	res := service.photoRepository.InsertPhoto(photo, id)
	return res
}

func (service *photoService) Get(id int) string {
	res := service.photoRepository.GetPhoto(id)
	base := helpers.ImageBase(res)
	return base
}

func (service *photoService) Delete(id int) models.Profile {
	res := service.photoRepository.DeletePhoto(id)
	return res
}
