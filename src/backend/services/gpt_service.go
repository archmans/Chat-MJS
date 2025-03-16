package services

import (
	"backend/config"
	"backend/models"
	"backend/algorithm"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProcessQuestion(c *gin.Context) {
	var request struct {
		Pertanyaan string `json:"pertanyaan"`
		Method     string `json:"method"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	fmt.Println("DEBUG - ProcessQuestion - pertanyaan:", request.Pertanyaan)
	fmt.Println("DEBUG - ProcessQuestion - method:", request.Method)

	var dbQuestions []models.ChatGPT
	config.DB.Find(&dbQuestions)

	var questionStrings []string
	for _, q := range dbQuestions {
		questionStrings = append(questionStrings, q.Pertanyaan)
	}

	fmt.Println("DEBUG - ProcessQuestion - listOfQuestions:", questionStrings)

	response := algorithm.ParseInput(request.Pertanyaan, dbQuestions, request.Method)

	fmt.Println("DEBUG - ProcessQuestion - response:", response)

	if response[0] == "error" {
		c.JSON(http.StatusBadRequest, gin.H{"response": response[1]})
		return
	}

	if response[0] == "kalender" || response[0] == "kalkulator" {
		c.JSON(http.StatusOK, gin.H{"response": response[1]})
		return
	}

	if response[0] == "tambah" {
		newQuestion := models.ChatGPT{
			Pertanyaan: response[1],
			Jawaban:    response[2],
		}

		var existing models.ChatGPT
		if config.DB.Where("pertanyaan = ?", response[1]).First(&existing).RowsAffected != 0 {
			config.DB.Model(&existing).Update("jawaban", response[2])
			c.JSON(http.StatusOK, gin.H{"response": "Pertanyaan diperbarui"})
		} else {
			config.DB.Create(&newQuestion)
			c.JSON(http.StatusOK, gin.H{"response": "Pertanyaan ditambahkan"})
		}
		return
	}

	if response[0] == "hapus" {
		var existing models.ChatGPT
		if config.DB.Where("pertanyaan = ?", response[1]).First(&existing).RowsAffected != 0 {
			config.DB.Delete(&existing)
			c.JSON(http.StatusOK, gin.H{"response": "Pertanyaan dihapus"})
		} else {
			c.JSON(http.StatusOK, gin.H{"response": "Pertanyaan tidak ditemukan"})
		}
		return
	}

	if response[0] == "jawaban" {
		c.JSON(http.StatusOK, gin.H{"response": response[1]})
		return
	}

	if response[0] == "rekomendasi" {
		c.JSON(http.StatusOK, gin.H{"response": response[1]})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "Jawaban tidak ditemukan"})
}
