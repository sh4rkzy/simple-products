package controllers

import (
	"github.com/gin-gonic/gin"
	uuids "github.com/sh4rkzy/infrastructure/utils"
	"net/http"
	"time"
)

func GetProducts(c *gin.Context) {
	response := gin.H{
		"status_code": 200,
		"message":     "OK",
		"products": []gin.H{
			{
				"product_id": "1",
			},
		},
		"transaction": gin.H{
			"transaction_id": uuids.GenerateUuid(),
			"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	c.JSON(http.StatusOK, response)
}
