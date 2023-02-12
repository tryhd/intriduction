package routers

import (
	"intoduction/app/controllers"
	"intoduction/app/repositories"
	"intoduction/app/services"
	"intoduction/configs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = configs.SetupDatabaseConnection()

	photoRepository repositories.PhotoRepository = repositories.NewPhotoRepository(db)

	photoService services.PhotoService = services.NewPhotoService(photoRepository)

	photoController controllers.PhotoController = controllers.NewPhotoController(photoService)
)

func PhotoRoute(route *gin.Engine) {

	ProfileRoutes := route.Group("api/v1/photo")
	{
		ProfileRoutes.GET("/:profile_code", photoController.Get)
		ProfileRoutes.DELETE("/:profile_code", photoController.Delete)
		ProfileRoutes.PUT("/:profile_code", photoController.Insert)
	}
}
