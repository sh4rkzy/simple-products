package products

import (
	"github.com/gin-gonic/gin"
	"github.com/sh4rkzy/infrastructure/database/repository"
	"github.com/sh4rkzy/infrastructure/utils"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

type ProductController struct {
	ProductRepo *repository.ProductRepository
}

func NewProductController(repo *repository.ProductRepository) *ProductController {
	return &ProductController{
		ProductRepo: repo,
	}
}


// Criar Produto
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var req struct {
		Name  string  `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Dados inválidos",
			"error":       err.Error(),
			"transaction": gin.H{
				"transaction_id": utils.GenerateUuid(),
				"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
			},
		})
		return
	}

	product := repository.Product{
		ID:    utils.GenerateUuid(),
		Name:  req.Name,
		Price: req.Price,
	}

	id, err := pc.ProductRepo.CreateProduct(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Erro ao criar produto",
			"error":       err.Error(),
			"transaction": gin.H{
				"transaction_id": utils.GenerateUuid(),
				"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status_code": http.StatusCreated,
		"message":     "Produto criado com sucesso",
		"product_id":  id,
		"transaction": gin.H{
			"transaction_id": utils.GenerateUuid(),
			"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
		},
	})
}

// Listar Produtos
func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.ProductRepo.GetProducts(c.Request.Context(), bson.M{}) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Erro ao buscar produtos",
			"error":       err.Error(),
			"transaction": gin.H{
				"transaction_id": utils.GenerateUuid(),
				"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
			},
		})
		return
	}

	var productsResponse []gin.H
	for _, product := range products {
		productsResponse = append(productsResponse, gin.H{
			"product_id": product.ID,
			"name":       product.Name,
			"price":      product.Price,
			"dt_created": product.DtCreated,
			"dt_updated": product.DtUpdated,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "OK",
		"products":    productsResponse,
		"transaction": gin.H{
			"transaction_id": utils.GenerateUuid(),
			"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
		},
	})
}