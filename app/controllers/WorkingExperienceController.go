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

type WorkingExperienceController interface {
	FindByID(context *gin.Context)

	Insert(context *gin.Context)

	Update(context *gin.Context)
}

type workingExperienceController struct {
	workingExperienceService services.WorkingExperienceService
}

func NewWorkingExperienceController(workingExperienceServ services.WorkingExperienceService) WorkingExperienceController {
	return &workingExperienceController{
		workingExperienceService: workingExperienceServ,
	}
}

func (c *workingExperienceController) FindByID(context *gin.Context) {
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	var workingExperience models.WorkingExperience = c.workingExperienceService.FindByID(intId)
	if (workingExperience == models.WorkingExperience{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		// WorkingExperienceUpdateDTO.WorkingExperienceCode = intId
		res := helpers.BuildResponse(true, "OK", workingExperience)
		context.JSON(http.StatusOK, res)
	}
}

func (c *workingExperienceController) Insert(context *gin.Context) {
	var workingExperienceCreateDTO dtos.WorkingExperienceCreateDTO
	context.ShouldBind(&workingExperienceCreateDTO)
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	workingExperienceCreateDTO.ProfileID = intId
	result := c.workingExperienceService.Insert(workingExperienceCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *workingExperienceController) Update(context *gin.Context) {
	var workingExperienceUpdateDTO dtos.WorkingExperienceUpdateDTO
	errDTO := context.ShouldBind(&workingExperienceUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	var workingExperience models.WorkingExperience = c.workingExperienceService.FindByID(intId)
	if (workingExperience == models.WorkingExperience{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		workingExperienceUpdateDTO.ProfileID = intId
		result := c.workingExperienceService.Update(workingExperienceUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}
