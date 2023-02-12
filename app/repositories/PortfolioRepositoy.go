package repositories

import (
	// "github.com/ydhnwb/golang_api/entity"

	"intoduction/app/models"

	"gorm.io/gorm"
)

type PortfolioRepository interface {
	InsertPortfolio(pf models.Portfolio) models.Portfolio
	UpdatePortfolio(pf models.Portfolio) models.Portfolio
	DeletePortfolio(pf models.Portfolio) models.Portfolio
	AllPortfolio() []models.Portfolio
	FindPortfolioByID(portfolioID string) models.Portfolio
	DeletedPortfolio() []models.Portfolio
	RestorePortfolio(pf models.Portfolio) models.Portfolio
	DeletePermanentPortfolio(pf models.Portfolio)
}

type portfolioConnection struct {
	connection *gorm.DB
}

func NewPortfolioRepository(dbConn *gorm.DB) PortfolioRepository {
	return &portfolioConnection{
		connection: dbConn,
	}
}

func (db *portfolioConnection) InsertPortfolio(portfolio models.Portfolio) models.Portfolio {
	db.connection.Create(&portfolio)
	db.connection.Find(&portfolio)
	return portfolio
}

func (db *portfolioConnection) UpdatePortfolio(portfolio models.Portfolio) models.Portfolio {
	db.connection.Save(&portfolio)
	db.connection.Find(&portfolio)
	return portfolio
}

func (db *portfolioConnection) DeletePortfolio(portfolio models.Portfolio) models.Portfolio {
	db.connection.Delete(&portfolio)
	db.connection.Find(&portfolio)
	return portfolio
}

func (db *portfolioConnection) AllPortfolio() []models.Portfolio {
	var portfolios []models.Portfolio
	db.connection.Find(&portfolios)
	return portfolios
}

func (db *portfolioConnection) FindPortfolioByID(portfolioID string) models.Portfolio {
	var portfolio models.Portfolio
	db.connection.Where("id =?", portfolioID).First(&portfolio)
	return portfolio
}

func (db *portfolioConnection) DeletedPortfolio() []models.Portfolio {
	var portfolios []models.Portfolio
	db.connection.Unscoped().Where("deleted_at != 0").Find(&portfolios)
	return portfolios
}

func (db *portfolioConnection) RestorePortfolio(portfolio models.Portfolio) models.Portfolio {
	db.connection.Unscoped().First(&portfolio).Update("deleted_at", nil)
	db.connection.Find(&portfolio)
	return portfolio
}

func (db *portfolioConnection) DeletePermanentPortfolio(portfolio models.Portfolio) {
	db.connection.Unscoped().Delete(&portfolio)
}
