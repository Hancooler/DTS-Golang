// routes.go

package routes

import (
	"perpus/controller"

	"github.com/gin-gonic/gin"
)

// SetupAPIRoutes menginisialisasi router dan mendefinisikan rute-rute untuk API RESTful
func SetupAPIRoutes(r *gin.Engine, bookController *controller.BookController, borrowingController *controller.BorrowingController, memberController *controller.MemberController, petugasController *controller.PetugasController) {
	// Rute untuk model Book
	books := r.Group("/api/books")
	{
		books.GET("/", bookController.FindAll)
		books.GET("/:id", bookController.Find)
		books.POST("/", bookController.Create)
		books.PUT("/:id", bookController.Update)
		books.DELETE("/:id", bookController.Delete)
	}

	// Rute untuk model Borrowing
	borrowings := r.Group("/api/borrowings")
	{
		borrowings.GET("/", borrowingController.FindAll)
		borrowings.GET("/:id", borrowingController.Find)
		borrowings.POST("/", borrowingController.Create)
		borrowings.PUT("/:id", borrowingController.Update)
		borrowings.DELETE("/:id", borrowingController.Delete)
	}

	// Rute untuk model Member
	members := r.Group("/api/members")
	{
		members.GET("/", memberController.FindAll)
		members.GET("/:id", memberController.Find)
		members.POST("/", memberController.Create)
		members.PUT("/:id", memberController.Update)
		members.DELETE("/:id", memberController.Delete)
	}

	// Rute untuk model Petugas
	petugas := r.Group("/api/petugas")
	{
		petugas.GET("/", petugasController.FindAll)
		petugas.GET("/:id", petugasController.Find)
		petugas.POST("/", petugasController.Create)
		petugas.PUT("/:id", petugasController.Update)
		petugas.DELETE("/:id", petugasController.Delete)
	}
}
