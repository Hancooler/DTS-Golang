package models

import (
	"perpus/controller"

	"github.com/gin-gonic/gin"
)

// SetupMemberRouter menginisialisasi router dan mendefinisikan rute-rute untuk model Member
func SetupMemberRouter(r *gin.Engine, memberController *controller.MemberController) {
	memberRoutes := r.Group("/members")

	// Rute untuk menampilkan semua anggota
	memberRoutes.GET("/", memberController.FindAll)

	// Rute untuk menampilkan detail anggota berdasarkan ID
	memberRoutes.GET("/:id", memberController.Find)

	// Rute untuk menambahkan anggota baru
	memberRoutes.POST("/", memberController.Create)

	// Rute untuk mengupdate informasi anggota berdasarkan ID
	memberRoutes.PUT("/:id", memberController.Update)

	// Rute untuk menghapus anggota berdasarkan ID
	memberRoutes.DELETE("/:id", memberController.Delete)
}
