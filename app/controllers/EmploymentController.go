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
	Get(context *gin.Context)
	Insert(context *gin.Context)
	Delete(context *gin.Context)
}

type employmentController struct {
	employmentService services.EmploymentService
}

func NewEmploymentController(employmentServ services.EmploymentService) EmploymentController {
	return &employmentController{
		employmentService: employmentServ,
	}
}

func (c *employmentController) Get(context *gin.Context) {
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	var employments []models.Employment = c.employmentService.Get(intId)

	res := helpers.BuildResponse(true, "OK", employments)
	context.JSON(http.StatusOK, res)
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

func (c *employmentController) Delete(context *gin.Context) {
	var employment models.Employment
	employmentId := context.Query("id")
	intEmploymentId, _ := strconv.Atoi(employmentId)

	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)

	employment.ID = intEmploymentId
	employment.ProfileID = intId
	result := c.employmentService.Delete(employment)
	res := helpers.BuildResponse(true, "Deleted", result)
	context.JSON(http.StatusOK, res)

}
