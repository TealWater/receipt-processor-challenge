package main

import (
	"net/http"

	"github.com/TealWater/fetch-rewards/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hi Mom!",
		})
	})

	receipts := r.Group("/receipts")
	{
		receipts.POST("/process", controller.ProcessReceipt)
		receipts.GET("/:id/points")
	}

	r.Run(":8080")
}
