package controller

import (
	"net/http"

	"perpus/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BorrowingController struct {
	DB *gorm.DB
}

func (controller *BorrowingController) Create(ctx *gin.Context) {
	var borrowing models.Borrowing
	if err := ctx.ShouldBindJSON(&borrowing); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.DB.Create(&borrowing)
	ctx.JSON(http.StatusCreated, borrowing)
}

func (controller *BorrowingController) Find(ctx *gin.Context) {
	var borrowing models.Borrowing
	if err := controller.DB.First(&borrowing, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	ctx.JSON(http.StatusOK, borrowing)
}

func (controller *BorrowingController) Update(ctx *gin.Context) {
	var borrowing models.Borrowing
	if err := controller.DB.First(&borrowing, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	if err := ctx.ShouldBindJSON(&borrowing); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.DB.Save(&borrowing)
	ctx.JSON(http.StatusOK, borrowing)
}

func (controller *BorrowingController) Delete(ctx *gin.Context) {
	var borrowing models.Borrowing
	if err := controller.DB.First(&borrowing, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	controller.DB.Delete(&borrowing)
	ctx.JSON(http.StatusNoContent, nil)
}

func (controller *BorrowingController) FindAll(ctx *gin.Context) {
	var borrowings []models.Borrowing
	controller.DB.Find(&borrowings)
	ctx.JSON(http.StatusOK, borrowings)
}
