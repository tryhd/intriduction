package controllers

import (
	"intoduction/app/dtos"
	"intoduction/app/helpers"
	"intoduction/app/models"
	"intoduction/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type contentController struct {
	contentService services.ContentService
}

func NewContentController(contentServ services.ContentService) ContentController {
	return &contentController{
		contentService: contentServ,
	}
}

func (c *contentController) All(context *gin.Context) {
	var contents []models.Content = c.contentService.All()
	res := helpers.BuildResponse(true, "OK", contents)
	context.JSON(http.StatusOK, res)
}

func (c *contentController) FindByID(context *gin.Context) {
	id := context.Param("id")
	var content models.Content = c.contentService.FindByID(id)
	if (content == models.Content{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", content)
		context.JSON(http.StatusOK, res)
	}
}

func (c *contentController) Insert(context *gin.Context) {
	var contentCreateDTO dtos.ContentCreateDTO
	imageName := helpers.FileUpload(context, "./public/content/", "content_image")
	contentCreateDTO.ContentImage = imageName
	errDTO := context.ShouldBind(&contentCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.contentService.Insert(contentCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *contentController) Update(context *gin.Context) {
	var contentUpdateDTO dtos.ContentUpdateDTO
	errDTO := context.ShouldBind(&contentUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	id := context.Param("id")
	var content models.Content = c.contentService.FindByID(id)
	if (content == models.Content{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		if context.Param("content_image") != "" {
			imageName := helpers.FileUpload(context, "./public/content/", "content_image")
			contentUpdateDTO.ContentImage = imageName
		} else {
			contentUpdateDTO.ContentImage = content.ContentImage
		}
		contentUpdateDTO.ID = id
		result := c.contentService.Update(contentUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *contentController) Delete(context *gin.Context) {
	var content models.Content
	id := context.Param("id")
	content = c.contentService.FindByID(id)
	if (content == models.Content{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		content.ID = id
		result := c.contentService.Delete(content)
		res := helpers.BuildResponse(true, "Deleted", result)
		context.JSON(http.StatusOK, res)
	}
}
