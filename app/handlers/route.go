package handlers

import (
	"intoduction/configs"
	"intoduction/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func Routing() {
	r := gin.New()
	r.Use(configs.CORSMiddleware())

	routers.ContactRoute(r) //Added all contact routes
	routers.SkillRoute(r)   //Added all skill routes
	routers.ProfileRoute(r) //Added all skill routes

	r.Run(":" + os.Getenv("APP_PORT"))
}
