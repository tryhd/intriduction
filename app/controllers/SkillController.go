package controllers

import (
	"intoduction/app/dtos"
	"intoduction/app/helpers"
	"intoduction/app/models"
	"intoduction/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SkillController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Delete(context *gin.Context)
}

type skillController struct {
	skillService services.SkillService
}

func NewSkillController(skillServ services.SkillService) SkillController {
	return &skillController{
		skillService: skillServ,
	}
}

func (c *skillController) All(context *gin.Context) {
	var skills []models.Skill = c.skillService.All()
	res := helpers.BuildResponse(true, "OK", skills)
	context.JSON(http.StatusOK, res)
}

func (c *skillController) Insert(context *gin.Context) {
	var skillCreateDTO dtos.SkillCreateDTO
	errDTO := context.ShouldBind(&skillCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.skillService.Insert(skillCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *skillController) Delete(context *gin.Context) {
	var skill models.Skill
	id := context.Param("id")
	skill = c.skillService.FindByID(id)
	if (skill == models.Skill{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		skill.ID = id
		result := c.skillService.Delete(skill)
		res := helpers.BuildResponse(true, "Deleted", result)
		context.JSON(http.StatusOK, res)
	}
}
