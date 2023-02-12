package services

import (
	"intoduction/app/dtos"
	"intoduction/app/models"
	"intoduction/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type PortfolioService interface {
	Insert(portfolio dtos.PortfolioCreateDTO) models.Portfolio
	Update(portfolio dtos.PortfolioUpdateDTO) models.Portfolio
	Restore(portfolio dtos.PortfolioRestoreDTO) models.Portfolio
	Delete(portfolio models.Portfolio) models.Portfolio
	DeletePermanent(portfolio models.Portfolio)
	All() []models.Portfolio
	Deleted() []models.Portfolio
	FindByID(portfolioID string) models.Portfolio
}

type portfolioService struct {
	portfolioRepository repositories.PortfolioRepository
}

func NewPortfolioService(portfolioRepo repositories.PortfolioRepository) PortfolioService {
	return &portfolioService{
		portfolioRepository: portfolioRepo,
	}
}

func (service *portfolioService) Insert(portfolio dtos.PortfolioCreateDTO) models.Portfolio {
	newPortfolio := models.Portfolio{}
	err := smapping.FillStruct(&newPortfolio, smapping.MapFields(&portfolio))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.portfolioRepository.InsertPortfolio(newPortfolio)
	return res
}

func (service *portfolioService) Update(portfolio dtos.PortfolioUpdateDTO) models.Portfolio {
	newPortfolio := models.Portfolio{}
	err := smapping.FillStruct(&newPortfolio, smapping.MapFields(&portfolio))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.portfolioRepository.UpdatePortfolio(newPortfolio)
	return res
}

func (service *portfolioService) Delete(portfolio models.Portfolio) models.Portfolio {
	return service.portfolioRepository.DeletePortfolio(portfolio)
}

func (service *portfolioService) All() []models.Portfolio {
	return service.portfolioRepository.AllPortfolio()
}

func (service *portfolioService) FindByID(portfolioID string) models.Portfolio {
	return service.portfolioRepository.FindPortfolioByID(portfolioID)
}

func (service *portfolioService) Deleted() []models.Portfolio {
	return service.portfolioRepository.DeletedPortfolio()
}

func (service *portfolioService) Restore(portfolio dtos.PortfolioRestoreDTO) models.Portfolio {
	newPortfolio := models.Portfolio{}
	err := smapping.FillStruct(&newPortfolio, smapping.MapFields(&portfolio))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.portfolioRepository.RestorePortfolio(newPortfolio)
	return res
}

func (service *portfolioService) DeletePermanent(portfolio models.Portfolio) {
	service.portfolioRepository.DeletePermanentPortfolio(portfolio)
}
