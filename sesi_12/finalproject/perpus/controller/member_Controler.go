package controller

import (
	"net/http"

	"perpus/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MemberController struct {
	DB *gorm.DB
}

func (controller *MemberController) Create(ctx *gin.Context) {
	var member models.Member
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.DB.Create(&member)
	ctx.JSON(http.StatusCreated, member)
}

func (controller *MemberController) Find(ctx *gin.Context) {
	var member models.Member
	if err := controller.DB.First(&member, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	ctx.JSON(http.StatusOK, member)
}

func (controller *MemberController) Update(ctx *gin.Context) {
	var member models.Member
	if err := controller.DB.First(&member, ctx.Params.ByName("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found "})
		return

	}
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.DB.Save(&member)
	ctx.JSON(http.StatusOK, member)

}

func (controller *MemberController) Delete(ctx *gin.Context) {
	var member models.Member
	if err := controller.DB.First(&member, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	controller.DB.Delete(&member)
	ctx.JSON(http.StatusNoContent, nil)
}

func (controller *MemberController) FindAll(ctx *gin.Context) {
	var members []models.Member
	controller.DB.Find(&members)
	ctx.JSON(http.StatusOK, members)
}
