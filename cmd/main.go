package main

import (
	"os"

	"github.com/akhmad-ardi/be-evowish/config"
	"github.com/akhmad-ardi/be-evowish/database"
	"github.com/akhmad-ardi/be-evowish/models"
	"github.com/akhmad-ardi/be-evowish/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":" + os.Getenv("PORT"))
}
