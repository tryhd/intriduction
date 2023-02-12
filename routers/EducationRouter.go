package routers

import (
	"intoduction/app/controllers"
	"intoduction/app/repositories"
	"intoduction/app/services"

	"github.com/gin-gonic/gin"
)

var (
	educationRepository repositories.EducationRepository = repositories.NewEducationRepository(db)

	educationService services.EducationService = services.NewEducationService(educationRepository)

	educationController controllers.EducationController = controllers.NewEducationController(educationService)
)

func EducationRoute(route *gin.Engine) {

	EducationRoutes := route.Group("api/v1/education")
	{
		EducationRoutes.GET("/:profile_code", educationController.Get)
		EducationRoutes.POST("/:profile_code", educationController.Insert)
		EducationRoutes.DELETE("/:profile_code", educationController.Delete)
	}
}
