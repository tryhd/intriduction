package routers

import (
	"intoduction/app/controllers"
	"intoduction/app/repositories"
	"intoduction/app/services"

	"github.com/gin-gonic/gin"
)

var (
	contactRepository repositories.ContactRepository = repositories.NewContactRepository(db)

	contactService services.ContactService = services.NewContactService(contactRepository)

	contactController controllers.ContactController = controllers.NewContactController(contactService)
)

func ContactRoute(route *gin.Engine) {

	contactRoutes := route.Group("api/v1/contact")
	{
		contactRoutes.GET("/", contactController.All)
		contactRoutes.POST("/", contactController.Insert)
		contactRoutes.GET("/:id", contactController.FindByID)
		contactRoutes.PUT("/:id", contactController.Update)
		contactRoutes.DELETE("/:id", contactController.Delete)
	}
}
