package main

import (
	"sesi_6/pkg/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":5000"
	func() *gin.Engine {
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.String(200, "Hello World")
		})
		api := r.Group("/api")
		employee := api.Group("/employees")
		employee.GET("", controller.GetEmployeeList)
		return r
	}().Run(port)
}
