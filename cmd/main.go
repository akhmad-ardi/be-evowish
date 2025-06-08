package main

import (
	"os"

	"github.com/akhmad-ardi/be-evowish/config"
	"github.com/akhmad-ardi/be-evowish/database"
	"github.com/akhmad-ardi/be-evowish/models"
	"github.com/akhmad-ardi/be-evowish/routes"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/akhmad-ardi/be-evowish/docs"
	swaggerFiles "github.com/swaggo/files"
)

// @title Evowish API
// @version 1.0
// @description Backend API untuk platform undangan digital multievent (Evowish).
// @host localhost:8080
// @BasePath /
func main() {
	config.LoadEnv()
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.SetupRoutes(r)
	r.Run(":" + os.Getenv("PORT"))
}
