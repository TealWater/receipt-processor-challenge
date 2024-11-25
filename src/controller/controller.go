package controller

import (
	"net/http"

	"github.com/TealWater/fetch-rewards/model"
	"github.com/TealWater/fetch-rewards/utility"
	"github.com/gin-gonic/gin"
)

var mps map[model.Id]model.Points

func ProcessReceipt(ctx *gin.Context) {
	receipt := &model.Receipt{}
	points := 0

	err := ctx.ShouldBindBodyWithJSON(receipt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pt, err := utility.ValidateName(receipt.Retailer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	points += pt

	pt, err = utility.ValidateTotal(receipt.Total)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	points += pt

	pt, err = utility.ValidatePurchaseTime(receipt.PurchaseTime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	points += pt

	pt, err = utility.ValidatePurchaseDate(receipt.PurchaseDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	points += pt

	pt, err = utility.ValidateItemDescription(receipt.Items)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	points += pt

	points += utility.CountItems(receipt.Items)

	ctx.JSON(http.StatusOK, model.Points{Points: points})
}
