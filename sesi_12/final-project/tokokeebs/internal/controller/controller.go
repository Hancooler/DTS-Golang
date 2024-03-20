package controller

import (
	"net/http"
	"toko-keyboard/internal/model"
	"toko-keyboard/internal/repository"

	"github.com/gin-gonic/gin"
)

func CreateKeyboard(c *gin.Context) {
	var keyboard model.Keyboard
	if err := c.ShouldBindJSON(&keyboard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.CreateKeyboard(&keyboard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create keyboard"})
		return
	}

	c.JSON(http.StatusCreated, keyboard)
}

func ListKeyboards(c *gin.Context) {
	keyboards, err := repository.ListKeyboards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch keyboards"})
		return
	}

	c.JSON(http.StatusOK, keyboards)
}

// Implementasi endpoint lainnya...
