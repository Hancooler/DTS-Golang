package controller

import (
	"net/http"
	"sesi_10/jwt/database"
	"sesi_10/jwt/helpers"
	"sesi_10/jwt/models"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GenerateFromPassword(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
	})

}

func UserLogin(c *gin.Context) {
	//db := database.GetDB()

	c.JSON(http.StatusOK, "Hello World")
}
