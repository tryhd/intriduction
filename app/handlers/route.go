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

	routers.ProfileRoute(r)           //Added all skill routes
	routers.WorkingExperienceRoute(r) //Added all skill routes
	routers.SkillRoute(r)             //Added all skill routes

	r.Run(":" + os.Getenv("APP_PORT"))
}
