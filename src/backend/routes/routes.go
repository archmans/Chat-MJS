package routes

import (
	"backend/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/gpt", services.ProcessQuestion)
	}
}
