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

type ProfileController interface {
	FindByID(context *gin.Context)

	Insert(context *gin.Context)

	Update(context *gin.Context)
}

type profileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileServ services.ProfileService) ProfileController {
	return &profileController{
		profileService: profileServ,
	}
}

func (c *profileController) FindByID(context *gin.Context) {
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	var Profile models.Profile = c.profileService.FindByID(intId)
	if Profile.ProfileCode == 0 {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		// profileUpdateDTO.ProfileCode = intId
		res := helpers.BuildResponse(true, "OK", Profile)
		context.JSON(http.StatusOK, res)
	}
}

func (c *profileController) Insert(context *gin.Context) {
	var profileCreateDTO dtos.ProfileCreateDTO
	context.ShouldBind(&profileCreateDTO)

	result := c.profileService.Insert(profileCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *profileController) Update(context *gin.Context) {
	var profileUpdateDTO dtos.ProfileUpdateDTO
	errDTO := context.ShouldBind(&profileUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	var profile models.Profile = c.profileService.FindByID(intId)
	if profile.ProfileCode == 0 {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		profileUpdateDTO.ProfileCode = intId
		result := c.profileService.Update(profileUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}
