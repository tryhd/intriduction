package routers

import (
	"intoduction/app/controllers"
	"intoduction/app/repositories"
	"intoduction/app/services"

	"github.com/gin-gonic/gin"
)

var (
	employmentRepository repositories.EmploymentRepository = repositories.NewEmploymentRepository(db)

	employmentService services.EmploymentService = services.NewEmploymentService(employmentRepository)

	employmentController controllers.EmploymentController = controllers.NewEmploymentController(employmentService)
)

func EmploymentRoute(route *gin.Engine) {

	EmploymentRoutes := route.Group("api/v1/employment")
	{
		EmploymentRoutes.GET("/:profile_code", employmentController.FindByID)
		EmploymentRoutes.POST("/:profile_code", employmentController.Insert)
		EmploymentRoutes.PUT("/:profile_code", employmentController.Update)
	}
}
