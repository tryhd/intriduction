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

	clientRepository repositories.ClientRepository = repositories.NewClientRepository(db)

	clientService services.ClientService = services.NewClientService(clientRepository)

	clientController controllers.ClientController = controllers.NewClientController(clientService)
)

func ClientRoute(route *gin.Engine) {

	clientRoutes := route.Group("api/v1/client")
	{
		clientRoutes.GET("/", clientController.All)
		clientRoutes.GET("/:id", clientController.FindByID)
		clientRoutes.POST("/", clientController.Insert)
		clientRoutes.PUT("/:id", clientController.Update)
		clientRoutes.DELETE("/:id", clientController.Delete)
		clientRoutes.GET("/deleted", clientController.Deleted)
		clientRoutes.PUT("/restore/:id", clientController.Restore)
		clientRoutes.DELETE("/delete/:id", clientController.DeletePermanent)
	}
}
