package routes

import (
	"github.com/akhmad-ardi/be-evowish/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers)
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)
}
