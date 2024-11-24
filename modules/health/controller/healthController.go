package healthcheck

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sh4rkzy/infrastructure/database"
	uuids "github.com/sh4rkzy/infrastructure/utils"
)

func HealthChecked(c *gin.Context) {
	err := database.Connector().Ping(context.Background(), nil)

	var response gin.H
	if err != nil {
		response = gin.H{
			"status_code": 400,
			"application": "OK",
			"databases":   "NOK",
			"detailed":    err.Error(),
			"transaction": gin.H{
				"transaction_id": uuids.GenerateUuid(),
				"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
			},
		}
	} else {
		response = gin.H{
			"status_code": 200,
			"application": "OK",
			"databases":   "OK",
			"transaction": gin.H{
				"transaction_id": uuids.GenerateUuid(),
				"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
			},
		}
	}

	// Envia a resposta JSON
	c.JSON(http.StatusOK, response)
}
