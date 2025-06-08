package controllers

import (
	"net/http"

	"github.com/akhmad-ardi/be-evowish/database"
	"github.com/akhmad-ardi/be-evowish/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
