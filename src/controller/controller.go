package controller

import (
	"errors"
	"net/http"

	"github.com/TealWater/fetch-rewards/model"
	"github.com/TealWater/fetch-rewards/utility"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var mps map[uuid.UUID]model.Points

func init() {
	mps = make(map[uuid.UUID]model.Points, 0)
}

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
		return
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

	id := uuid.New()
	mps[id] = model.Points{Points: points}

	ctx.JSON(http.StatusOK, model.Id{ID: id.String()})
}

func FetchPoints(ctx *gin.Context) {
	if len(mps) == 0 {
		ctx.JSON(http.StatusNotFound, errors.New("error: ID does not exist in the database").Error())
		return
	}

	id, _ := uuid.Parse(ctx.Param("id"))
	val, ok := mps[id]
	if !ok {
		ctx.JSON(http.StatusNotFound, errors.New("error: ID does not exist in the database").Error())
		return
	}
	ctx.JSON(http.StatusOK, model.Points{Points: val.Points})
}
