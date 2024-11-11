package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sh4rkzy/infrastructure/database"
	"github.com/sh4rkzy/modules/health/controller"
	"github.com/sh4rkzy/modules/products/controllers"
)

func main() {
	ctx := context.Background()

	// Conecte-se ao banco de dados
	client := database.Connector()
	if err := client.Ping(ctx, nil); err != nil {
		fmt.Println("Erro ao conectar ao banco de dados")
		return
	}
	defer client.Disconnect(ctx)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/health", healthcheck.HealthChecked)
		api.GET("/products", products.GetProducts)
	}

	port := ":8080"
	fmt.Println("Server running on port", port)
	router.Run(port)
}
