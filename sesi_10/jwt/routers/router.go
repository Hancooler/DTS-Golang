package routers

import (
	"sesi_10/jwt/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")

	userRouter.POST("/register", controller.UserRegister)
	userRouter.GET("/login", controller.UserLogin)
	return r
}
