package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sh4rkzy/infrastructure/database"
	"github.com/sh4rkzy/infrastructure/database/repository"
	"github.com/sh4rkzy/modules/health/controller"
	"github.com/sh4rkzy/modules/products/controllers"
)

func main() {
	// Define um contexto com timeout de 5 segundos para a conexão com o banco de dados
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Tenta conectar ao banco de dados usando o contexto com timeout
	client := database.Connector()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Criação do repositório e controllers
	productRepo := repository.NewProductRepository(client, "appdb", "products")
	productController := products.NewProductController(productRepo)

	// Inicializa o servidor HTTP
	router := gin.Default()

	// Define as rotas da aplicação junto com os controllers
	api := router.Group("/api/v1")
	{
		api.GET("/health", healthcheck.HealthChecked)
		api.GET("/products", productController.GetProducts)
		api.POST("/products", productController.CreateProduct)
	}

	// Inicia o servidor na porta 8080
	port := ":3000"
	fmt.Println("Server running on port", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
