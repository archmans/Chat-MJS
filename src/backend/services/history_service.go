package services

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHistory(c *gin.Context) {
	var history []models.History
	config.DB.Find(&history)
	c.JSON(http.StatusOK, gin.H{"history": history})
}

func AddHistory(c *gin.Context) {
	var history models.History
	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&history)
	c.JSON(http.StatusOK, gin.H{"message": "History berhasil ditambahkan"})
}
