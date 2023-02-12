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

type SkillController interface {
	Get(context *gin.Context)
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

func (c *skillController) Get(context *gin.Context) {
	var skills []models.Skill = c.skillService.Get()

	res := helpers.BuildResponse(true, "OK", skills)
	context.JSON(http.StatusOK, res)
}

func (c *skillController) Insert(context *gin.Context) {
	var skillCreateDTO dtos.SkillCreateDTO
	errDTO := context.ShouldBind(&skillCreateDTO)
	id := context.Param("profile_code")
	intId, _ := strconv.Atoi(id)
	skillCreateDTO.ProfileID = intId
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
	id := context.Query("profile_code")
	intId, _ := strconv.Atoi(id)

	skill.ID = intId
	result := c.skillService.Delete(skill)
	res := helpers.BuildResponse(true, "Deleted", result)
	context.JSON(http.StatusOK, res)

}
