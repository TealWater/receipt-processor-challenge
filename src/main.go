package main

import (
	"net/http"
	"time"

	"github.com/TealWater/fetch-rewards/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"server time": time.DateTime,
		})
	})

	receipts := r.Group("/receipts")
	{
		receipts.POST("/process", controller.ProcessReceipt)
		receipts.GET("/:id/points", controller.FetchPoints)
	}

	r.Run(":8080")
}
