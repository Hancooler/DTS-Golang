package controller

import (
	"net/http"

	"perpus/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func (controller *BookController) Create(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.DB.Create(&book)
	ctx.JSON(http.StatusCreated, book)
}

func (controller *BookController) Find(ctx *gin.Context) {
	var book models.Book
	if err := controller.DB.First(&book, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (controller *BookController) Update(ctx *gin.Context) {
	var book models.Book
	if err := controller.DB.First(&book, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.DB.Save(&book)
	ctx.JSON(http.StatusOK, book)
}

func (controller *BookController) Delete(ctx *gin.Context) {
	var book models.Book
	if err := controller.DB.First(&book, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	controller.DB.Delete(&book)
	ctx.JSON(http.StatusNoContent, nil)
}

func (controller *BookController) FindAll(ctx *gin.Context) {
	var books []models.Book
	controller.DB.Find(&books)
	ctx.JSON(http.StatusOK, books)
}
