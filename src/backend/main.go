package main

import (
	"backend/config"
	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	r.Use(cors.Default())

	routes.SetupRoutes(r)

	r.Run(":8000")
}