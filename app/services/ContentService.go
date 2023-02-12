package services

import (
	"intoduction/app/dtos"
	"intoduction/app/models"
	"intoduction/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type ContentService interface {
	Insert(content dtos.ContentCreateDTO) models.Content
	Update(content dtos.ContentUpdateDTO) models.Content
	Delete(content models.Content) models.Content
	All() []models.Content
	FindByID(contentID string) models.Content
}

type contentService struct {
	contentRepository repositories.ContentRepository
}

func NewContentService(contentRepo repositories.ContentRepository) ContentService {
	return &contentService{
		contentRepository: contentRepo,
	}
}

func (service *contentService) Insert(content dtos.ContentCreateDTO) models.Content {
	newContent := models.Content{}
	err := smapping.FillStruct(&newContent, smapping.MapFields(&content))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.contentRepository.InsertContent(newContent)
	return res
}

func (service *contentService) Update(content dtos.ContentUpdateDTO) models.Content {
	newcontent := models.Content{}
	err := smapping.FillStruct(&newcontent, smapping.MapFields(&content))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.contentRepository.UpdateContent(newcontent)
	return res
}

func (service *contentService) Delete(content models.Content) models.Content {
	return service.contentRepository.DeleteContent(content)
}

func (service *contentService) All() []models.Content {
	return service.contentRepository.AllContent()
}

func (service *contentService) FindByID(contentID string) models.Content {
	return service.contentRepository.FindContentByID(contentID)
}
