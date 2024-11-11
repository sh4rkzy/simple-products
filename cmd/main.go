package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sh4rkzy/infrastructure/database"
	"github.com/sh4rkzy/modules/health/controller"
	"github.com/sh4rkzy/modules/products/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/health", healthController.HealthChecked)
	router.GET("/products", controllers.GetProducts)
	port := ":8080"
	fmt.Println("Server running on port", port)

	if err := database.Connector(); err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	router.Run(port)
}
