package controller

import (
	"net/http"
	"perpus/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	db *gorm.DB
}

func (c *BookController) Create(ctx *gin.Context) {
	var book models.Book
	ctx.BindJSON(&book)

	if err := c.db.Create(&book).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func (c *BookController) Update(ctx *gin.Context) {
	var book models.Book
	id := ctx.Param("id")

	if err := c.db.Where("id = ?", id).First(&book).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.BindJSON(&book)

	if err := c.db.Save(&book).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.db.Where("id = ?", id).Delete(&models.Book{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *BookController) Find(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if err := c.db.Where("id = ?", id).First(&book).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) FindAll(ctx *gin.Context) {
	var books []models.Book

	if err := c.db.Find(&books).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}
