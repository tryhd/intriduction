package routers

import (
	"intoduction/app/controllers"
	"intoduction/app/repositories"
	"intoduction/app/services"

	"github.com/gin-gonic/gin"
)

var (
	profileRepository repositories.ProfileRepository = repositories.NewProfileRepository(db)

	profileService services.ProfileService = services.NewProfileService(profileRepository)

	profileController controllers.ProfileController = controllers.NewProfileController(profileService)
)

func ProfileRoute(route *gin.Engine) {

	ProfileRoutes := route.Group("api/v1/profile")
	{
		ProfileRoutes.GET("/:profile_code", profileController.FindByID)
		ProfileRoutes.POST("/create", profileController.Insert)
		ProfileRoutes.PUT("/:profile_code", profileController.Update)
	}
}
