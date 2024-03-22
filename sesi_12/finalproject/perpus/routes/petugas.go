package models

import (
	"perpus/controller"

	"github.com/gin-gonic/gin"
)

// SetupPetugasRouter menginisialisasi router dan mendefinisikan rute-rute untuk model Petugas
func SetupPetugasRouter(r *gin.Engine, petugasController *controller.PetugasController) {
	petugasRoutes := r.Group("/petugas")

	// Rute untuk menampilkan semua petugas
	petugasRoutes.GET("/", petugasController.FindAll)

	// Rute untuk menampilkan detail petugas berdasarkan ID
	petugasRoutes.GET("/:id", petugasController.Find)

	// Rute untuk menambahkan petugas baru
	petugasRoutes.POST("/", petugasController.Create)

	// Rute untuk mengupdate informasi petugas berdasarkan ID
	petugasRoutes.PUT("/:id", petugasController.Update)

	// Rute untuk menghapus petugas berdasarkan ID
	petugasRoutes.DELETE("/:id", petugasController.Delete)
}
