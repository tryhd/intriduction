package controllers

import (
	"intoduction/app/dtos"
	"intoduction/app/helpers"
	"intoduction/app/models"
	"intoduction/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmploymentController interface {
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
}

type employmentController struct {
	employmentService services.EmploymentService
}

func NewEmploymentController(employmentServ services.EmploymentService) EmploymentController {
	return &employmentController{
		employmentService: employmentServ,
	}
}

func (c *employmentController) FindByID(context *gin.Context) {
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	var employment models.Employment = c.employmentService.FindByID(intId)
	if (employment == models.Employment{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		// EmploymentUpdateDTO.EmploymentCode = intId
		res := helpers.BuildResponse(true, "OK", employment)
		context.JSON(http.StatusOK, res)
	}
}

func (c *employmentController) Insert(context *gin.Context) {
	var employmentCreateDTO dtos.EmploymentCreateDTO
	context.ShouldBind(&employmentCreateDTO)
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	employmentCreateDTO.ProfileID = intId
	result := c.employmentService.Insert(employmentCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *employmentController) Update(context *gin.Context) {
	var employmentUpdateDTO dtos.EmploymentUpdateDTO
	errDTO := context.ShouldBind(&employmentUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	var employment models.Employment = c.employmentService.FindByID(intId)
	if (employment == models.Employment{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		employmentUpdateDTO.ProfileID = intId
		result := c.employmentService.Update(employmentUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}
