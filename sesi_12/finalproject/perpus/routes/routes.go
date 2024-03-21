package routes

import (
	"perpus/config"
	"perpus/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Inisialisasi koneksi database
	db, err := config.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Inisialisasi setiap controller dengan koneksi database yang telah disediakan
	bookController := &controller.BookController{db: db}
	borrowingController := &controller.BorrowingController{db: db}
	memberController := &controller.MemberController{db: db}
	petugasController := &controller.PetugasController{db: db}

	r.POST("/books", bookController.Create)
	r.GET("/books/:id", bookController.Find)
	r.PUT("/books/:id", bookController.Update)
	r.DELETE("/books/:id", bookController.Delete)
	r.GET("/books", bookController.FindAll)

	r.POST("/borrowings", borrowingController.Create)
	r.GET("/borrowings/:id", borrowingController.Find)
	r.PUT("/borrowings/:id", borrowingController.Update)
	r.DELETE("/borrowings/:id", borrowingController.Delete)
	r.GET("/borrowings", borrowingController.FindAll)

	r.POST("/members", memberController.Create)
	r.GET("/members/:id", memberController.Find)
	r.PUT("/members/:id", memberController.Update)
	r.DELETE("/members/:id", memberController.Delete)
	r.GET("/members", memberController.FindAll)

	r.POST("/petugas", petugasController.Create)
	r.GET("/petugas/:id", petugasController.Find)
	r.PUT("/petugas/:id", petugasController.Update)
	r.DELETE("/petugas/:id", petugasController.Delete)
	r.GET("/petugas", petugasController.FindAll)

	r.Run(":5000")
}
