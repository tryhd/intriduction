package controllers

import (
	"intoduction/app/dtos"
	"intoduction/app/helpers"
	"intoduction/app/models"
	"intoduction/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type contactController struct {
	contactService services.ContactService
}

func NewContactController(contactServ services.ContactService) ContactController {
	return &contactController{
		contactService: contactServ,
	}
}

func (c *contactController) All(context *gin.Context) {
	var contacts []models.Contact = c.contactService.All()
	res := helpers.BuildResponse(true, "OK", contacts)
	context.JSON(http.StatusOK, res)
}

func (c *contactController) FindByID(context *gin.Context) {
	id := context.Param("id")
	var contact models.Contact = c.contactService.FindByID(id)
	if (contact == models.Contact{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", contact)
		context.JSON(http.StatusOK, res)
	}
}

func (c *contactController) Insert(context *gin.Context) {
	var contactCreateDTO dtos.ContactCreateDTO
	errDTO := context.ShouldBind(&contactCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.contactService.Insert(contactCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *contactController) Update(context *gin.Context) {
	var contactUpdateDTO dtos.ContactUpdateDTO
	errDTO := context.ShouldBind(&contactUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	id := context.Param("id")
	var contact models.Contact = c.contactService.FindByID(id)
	if (contact == models.Contact{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		contactUpdateDTO.ID = id
		result := c.contactService.Update(contactUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *contactController) Delete(context *gin.Context) {
	var contact models.Contact
	id := context.Param("id")
	contact = c.contactService.FindByID(id)
	if (contact == models.Contact{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		contact.ID = id
		result := c.contactService.Delete(contact)
		res := helpers.BuildResponse(true, "Deleted", result)
		context.JSON(http.StatusOK, res)
	}
}
