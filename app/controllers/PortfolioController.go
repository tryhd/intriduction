package controllers

import (
	"intoduction/app/dtos"
	"intoduction/app/helpers"
	"intoduction/app/models"
	"intoduction/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PortfolioController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Deleted(context *gin.Context)
	Restore(context *gin.Context)
	DeletePermanent(context *gin.Context)
}

type portfolioController struct {
	portfolioService services.PortfolioService
}

func NewPortfolioController(portfolioServ services.PortfolioService) PortfolioController {
	return &portfolioController{
		portfolioService: portfolioServ,
	}
}

func (c *portfolioController) All(context *gin.Context) {
	var portfolios []models.Portfolio = c.portfolioService.All()
	res := helpers.BuildResponse(true, "OK", portfolios)
	context.JSON(http.StatusOK, res)
}

func (c *portfolioController) FindByID(context *gin.Context) {
	id := context.Param("id")
	var portfolio models.Portfolio = c.portfolioService.FindByID(id)
	if (portfolio == models.Portfolio{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", portfolio)
		context.JSON(http.StatusOK, res)
	}
}

func (c *portfolioController) Insert(context *gin.Context) {
	var portfolioCreateDTO dtos.PortfolioCreateDTO
	imageName := helpers.FileUpload(context, "/root/Go/src/intoduction/public/img/portfolio/", "portfolio_image")
	portfolioCreateDTO.PortfolioImage = imageName
	context.ShouldBind(&portfolioCreateDTO)

	result := c.portfolioService.Insert(portfolioCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *portfolioController) Update(context *gin.Context) {
	var portfolioUpdateDTO dtos.PortfolioUpdateDTO
	context.ShouldBind(&portfolioUpdateDTO)
	id := context.Param("id")
	var portfolio models.Portfolio = c.portfolioService.FindByID(id)
	if (portfolio == models.Portfolio{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		if context.Param("portfolio_image") != "" {
			imageName := helpers.FileUpload(context, "/root/Go/src/intoduction/public/img/portfolio/", "portfolio_image")
			portfolioUpdateDTO.PortfolioImage = imageName
		} else {
			portfolioUpdateDTO.PortfolioImage = portfolio.PortfolioImage
		}
		portfolioUpdateDTO.ID = id
		result := c.portfolioService.Update(portfolioUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *portfolioController) Delete(context *gin.Context) {
	var portfolio models.Portfolio
	id := context.Param("id")
	portfolio = c.portfolioService.FindByID(id)
	if (portfolio == models.Portfolio{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		portfolio.ID = id
		result := c.portfolioService.Delete(portfolio)
		res := helpers.BuildResponse(true, "Deleted", result)
		context.JSON(http.StatusOK, res)
	}
}

func (c *portfolioController) Deleted(context *gin.Context) {
	var portfolios []models.Portfolio = c.portfolioService.Deleted()
	res := helpers.BuildResponse(true, "OK", portfolios)
	context.JSON(http.StatusOK, res)
}

func (c *portfolioController) Restore(context *gin.Context) {
	var portfolioRestoreDTO dtos.PortfolioRestoreDTO
	id := context.Param("id")
	var portfolio models.Portfolio = c.portfolioService.FindByID(id)
	if (portfolio != models.Portfolio{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		portfolioRestoreDTO.ID = id
		result := c.portfolioService.Restore(portfolioRestoreDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *portfolioController) DeletePermanent(context *gin.Context) {
	var portfolio models.Portfolio
	id := context.Param("id")
	portfolio = c.portfolioService.FindByID(id)
	if (portfolio != models.Portfolio{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		portfolio.ID = id
		c.portfolioService.DeletePermanent(portfolio)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}
