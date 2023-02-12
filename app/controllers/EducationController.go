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

type EducationController interface {
	Get(context *gin.Context)
	Insert(context *gin.Context)
	Delete(context *gin.Context)
}

type educationController struct {
	educationService services.EducationService
}

func NewEducationController(educationServ services.EducationService) EducationController {
	return &educationController{
		educationService: educationServ,
	}
}

func (c *educationController) Get(context *gin.Context) {
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	var educations []models.Education = c.educationService.Get(intId)

	res := helpers.BuildResponse(true, "OK", educations)
	context.JSON(http.StatusOK, res)
}

func (c *educationController) Insert(context *gin.Context) {
	var educationCreateDTO dtos.EducationCreateDTO
	context.ShouldBind(&educationCreateDTO)
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	educationCreateDTO.ProfileID = intId
	result := c.educationService.Insert(educationCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *educationController) Delete(context *gin.Context) {
	var education models.Education
	educationId := context.Query("id")
	intEducationId, _ := strconv.Atoi(educationId)

	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)

	education.ID = intEducationId
	education.ProfileID = intId
	result := c.educationService.Delete(education)
	res := helpers.BuildResponse(true, "Deleted", result)
	context.JSON(http.StatusOK, res)

}
