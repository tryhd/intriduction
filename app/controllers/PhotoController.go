package controllers

import (
	"intoduction/app/dtos"
	"intoduction/app/helpers"
	"intoduction/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	Insert(context *gin.Context)
	Get(context *gin.Context)
	Delete(context *gin.Context)
}

type photoController struct {
	photoService services.PhotoService
}

func NewPhotoController(photoServ services.PhotoService) PhotoController {
	return &photoController{
		photoService: photoServ,
	}
}

func (c *photoController) Insert(context *gin.Context) {
	var photoCreateDTO dtos.PhotoCreateDTO
	errDTO := context.ShouldBind(&photoCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	imageName := helpers.FileUpload(id, photoCreateDTO.Base64img)

	result := c.photoService.Insert(imageName, intId)

	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *photoController) Get(context *gin.Context) {
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)

	result := c.photoService.Get(intId)

	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *photoController) Delete(context *gin.Context) {
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)

	result := c.photoService.Delete(intId)

	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}
