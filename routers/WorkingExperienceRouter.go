package routers

import (
	"intoduction/app/controllers"
	"intoduction/app/repositories"
	"intoduction/app/services"

	"github.com/gin-gonic/gin"
)

var (
	workingExperienceRepository repositories.WorkingExperienceRepository = repositories.NewWorkingExperienceRepository(db)

	workingExperienceService services.WorkingExperienceService = services.NewWorkingExperienceService(workingExperienceRepository)

	workingExperienceController controllers.WorkingExperienceController = controllers.NewWorkingExperienceController(workingExperienceService)
)

func WorkingExperienceRoute(route *gin.Engine) {

	WorkingExperienceRoutes := route.Group("api/v1/working-experience")
	{
		WorkingExperienceRoutes.GET("/:profile_code", workingExperienceController.FindByID)
		WorkingExperienceRoutes.POST("/:profile_code", workingExperienceController.Insert)
		WorkingExperienceRoutes.PUT("/:profile_code", workingExperienceController.Update)
	}
}
