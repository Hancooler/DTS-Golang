package main

import (
	"sesi_7/pkg/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	port := `:5000`
	func() *ginEngine {
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, "hello world")

		})
		api := r.Group("/api")
		employee := api.Group("/employees")
		employee.GET("", controller.GETEmployeeList)
		employee.POST("", controller.CreateEmployee)
		return r

	}().Run(port)

}
