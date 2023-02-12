package controllers

import (
	"intoduction/app/dtos"
	"intoduction/app/helpers"
	"intoduction/app/models"
	"intoduction/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Deleted(context *gin.Context)
	Restore(context *gin.Context)
	DeletePermanent(context *gin.Context)
}

type clientController struct {
	clientService services.ClientService
}

func NewClientController(clientServ services.ClientService) ClientController {
	return &clientController{
		clientService: clientServ,
	}
}

func (c *clientController) All(context *gin.Context) {
	var clients []models.Client = c.clientService.All()
	res := helpers.BuildResponse(true, "OK", clients)
	context.JSON(http.StatusOK, res)
}

func (c *clientController) FindByID(context *gin.Context) {
	id := context.Param("id")
	var client models.Client = c.clientService.FindByID(id)
	if (client == models.Client{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", client)
		context.JSON(http.StatusOK, res)
	}
}

func (c *clientController) Insert(context *gin.Context) {
	var clientCreateDTO dtos.ClientCreateDTO
	imageName := helpers.FileUpload(context, "/public/img/client/", "client_image")
	clientCreateDTO.ClientImage = imageName
	errDTO := context.ShouldBind(&clientCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.clientService.Insert(clientCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *clientController) Update(context *gin.Context) {
	var clientUpdateDTO dtos.ClientUpdateDTO
	errDTO := context.ShouldBind(&clientUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	id := context.Param("id")
	var client models.Client = c.clientService.FindByID(id)
	if (client == models.Client{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		if context.Param("client_image") != "" {
			imageName := helpers.FileUpload(context, "/public/img/client/", "client_image")
			clientUpdateDTO.ClientImage = imageName
		} else {
			clientUpdateDTO.ClientImage = client.ClientImage
		}
		clientUpdateDTO.ID = id
		result := c.clientService.Update(clientUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *clientController) Delete(context *gin.Context) {
	var client models.Client
	id := context.Param("id")
	client = c.clientService.FindByID(id)
	if (client == models.Client{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		client.ID = id
		result := c.clientService.Delete(client)
		res := helpers.BuildResponse(true, "Deleted", result)
		context.JSON(http.StatusOK, res)
	}
}

func (c *clientController) Deleted(context *gin.Context) {
	var clients []models.Client = c.clientService.Deleted()
	res := helpers.BuildResponse(true, "OK", clients)
	context.JSON(http.StatusOK, res)
}

func (c *clientController) Restore(context *gin.Context) {
	var clientRestoreDTO dtos.ClientRestoreDTO
	id := context.Param("id")
	var client models.Client = c.clientService.FindByID(id)
	if (client != models.Client{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		clientRestoreDTO.ID = id
		result := c.clientService.Restore(clientRestoreDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *clientController) DeletePermanent(context *gin.Context) {
	var client models.Client
	id := context.Param("id")
	client = c.clientService.FindByID(id)
	if (client != models.Client{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		client.ID = id
		c.clientService.DeletePermanent(client)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}
