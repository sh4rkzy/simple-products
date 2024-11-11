package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sh4rkzy/modules/health/controller"
	"github.com/sh4rkzy/modules/products/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/health", healthController.HealthChecked)
	router.GET("/products", controllers.GetProducts)
	port := ":8080"
	fmt.Println("Server running on port", port)
	router.Run(port)
}
