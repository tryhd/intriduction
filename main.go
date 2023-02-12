package main

import (
	"intoduction/configs"
	"intoduction/routers"
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = configs.SetupDatabaseConnection()

func main() {
	defer configs.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}))
	r.Static("/image/client", "/root/Go/src/intoduction/public/img/client")
	r.Static("/image/portfolio", "/root/Go/src/intoduction/public/img/portfolio")
	routers.ClientRoute(r)    //Added all client routes
	routers.ContactRoute(r)   //Added all contact routes
	routers.PortfolioRoute(r) //Added all portfolio routes
	routers.SkillRoute(r)     //Added all skill routes
	routers.ContentRoute(r)   //Added all skill routes
	routers.ProfileRoute(r)   //Added all skill routes

	r.Run(":" + os.Getenv("APP_PORT"))
}
