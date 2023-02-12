package routers

import (
	"intoduction/app/controllers"
	"intoduction/app/repositories"
	"intoduction/app/services"

	"github.com/gin-gonic/gin"
)

var (
	contentRepository repositories.ContentRepository = repositories.NewContentRepository(db)

	contentService services.ContentService = services.NewContentService(contentRepository)

	contentController controllers.ContentController = controllers.NewContentController(contentService)
)

func ContentRoute(route *gin.Engine) {

	contentRoutes := route.Group("api/v1/content")
	{
		contentRoutes.GET("/", contentController.All)
		contentRoutes.GET("/:id", contentController.FindByID)
		contentRoutes.POST("/", contentController.Insert)
		contentRoutes.PUT("/:id", contentController.Update)
		contentRoutes.DELETE("/:id", contentController.Delete)
	}
}
