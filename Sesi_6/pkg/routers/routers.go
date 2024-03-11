package routers

import (
	"sesi_6/pkg/controller"

	"github.com/gin-gonic/gin"
)

func Startserver(port string) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	api := r.Group("/api")

	employee := api.Group("/employees")
	employee.GET("", controller.GetEmployeeList)

	return r
}
