package routers

import (
	"intoduction/app/controllers"
	"intoduction/app/repositories"
	"intoduction/app/services"

	"github.com/gin-gonic/gin"
)

var (
	skillRepository repositories.SkillRepository = repositories.NewSkillRepository(db)

	skillService services.SkillService = services.NewSkillService(skillRepository)

	skillController controllers.SkillController = controllers.NewSkillController(skillService)
)

func SkillRoute(route *gin.Engine) {

	skillRoutes := route.Group("api/v1/skill")
	{
		skillRoutes.GET("/:profile_code", skillController.Get)
		skillRoutes.POST("/:profile_code", skillController.Insert)
		skillRoutes.DELETE("/:profile_code", skillController.Delete)
	}
}
