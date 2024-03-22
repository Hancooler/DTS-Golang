package models

import (
	"perpus/controller"

	"github.com/gin-gonic/gin"
)

// SetupBorrowingRouter menginisialisasi router dan mendefinisikan rute-rute untuk model Borrowing
func SetupBorrowingRouter(r *gin.Engine, borrowingController *controller.BorrowingController) {
	borrowingRoutes := r.Group("/borrowings")

	// Rute untuk menampilkan semua peminjaman
	borrowingRoutes.GET("/", borrowingController.FindAll)

	// Rute untuk menampilkan detail peminjaman berdasarkan ID
	borrowingRoutes.GET("/:id", borrowingController.Find)

	// Rute untuk menambahkan peminjaman baru
	borrowingRoutes.POST("/", borrowingController.Create)

	// Rute untuk mengupdate informasi peminjaman berdasarkan ID
	borrowingRoutes.PUT("/:id", borrowingController.Update)

	// Rute untuk menghapus peminjaman berdasarkan ID
	borrowingRoutes.DELETE("/:id", borrowingController.Delete)
}
