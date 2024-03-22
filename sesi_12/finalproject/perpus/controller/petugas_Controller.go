package controller

import (
	"net/http"
	"perpus/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PetugasController struct {
	DB *gorm.DB
}

func (controller *PetugasController) Create(ctx *gin.Context) {
	var petugas models.Petugas
	if err := ctx.ShouldBindJSON(&petugas); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.DB.Create(&petugas)
	ctx.JSON(http.StatusCreated, petugas)
}

func (controller *PetugasController) Find(ctx *gin.Context) {
	var petugas models.Petugas
	if err := controller.DB.First(&petugas, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	ctx.JSON(http.StatusOK, petugas)
}

func (controller *PetugasController) Update(ctx *gin.Context) {
	var petugas models.Petugas
	if err := controller.DB.First(&petugas, ctx.Params.ByName("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found "})
		return

	}
	if err := ctx.ShouldBindJSON(&petugas); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.DB.Save(&petugas)
	ctx.JSON(http.StatusOK, petugas)

}

func (controller *PetugasController) Delete(ctx *gin.Context) {
	var petugas models.Petugas
	if err := controller.DB.First(&petugas, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	controller.DB.Delete(&petugas)
	ctx.JSON(http.StatusNoContent, nil)
}

func (controller *PetugasController) FindAll(ctx *gin.Context) {
	var petugas []models.Petugas
	controller.DB.Find(&petugas)
	ctx.JSON(http.StatusOK, petugas)
}
