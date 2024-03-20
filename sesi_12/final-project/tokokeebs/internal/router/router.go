package router

import (
	"tokokeebs/internal/controller"
	"tokokeebs/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes (tanpa autentikasi)
	r.POST("/login", controller.Login)

	// Protected routes (membutuhkan autentikasi)
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	// Admin routes (memerlukan autentikasi sebagai admin)
	admin := protected.Group("/admin")
	admin.Use(middleware.AdminOnly())

	admin.POST("/keyboards", controller.CreateKeyboard)
	admin.GET("/keyboards", controller.ListKeyboards)
	admin.GET("/keyboards/:id", controller.GetKeyboard)
	admin.PUT("/keyboards/:id", controller.UpdateKeyboard)
	admin.DELETE("/keyboards/:id", controller.DeleteKeyboard)

	return r
}
