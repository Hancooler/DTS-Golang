package models

import (
	"perpus/controller"

	"github.com/gin-gonic/gin"
)

// SetupBookRouter menginisialisasi router dan mendefinisikan rute-rute untuk model Book
func SetupBookRouter(r *gin.Engine, bookController *controller.BookController) {
	bookRoutes := r.Group("/books")

	// Rute untuk menampilkan semua buku
	bookRoutes.GET("/", bookController.FindAll)

	// Rute untuk menampilkan detail buku berdasarkan ID
	bookRoutes.GET("/:id", bookController.Find)

	// Rute untuk menambahkan buku baru
	bookRoutes.POST("/", bookController.Create)

	// Rute untuk mengupdate informasi buku berdasarkan ID
	bookRoutes.PUT("/:id", bookController.Update)

	// Rute untuk menghapus buku berdasarkan ID
	bookRoutes.DELETE("/:id", bookController.Delete)
}
