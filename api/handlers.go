package api

import (
	currency "backend/concurrency"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertHandler(c *gin.Context) {
	var requestData ConversionRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amount, err := strconv.ParseFloat(requestData.Amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount must be a valid number"})
		return
	}

	result, err := currency.ConvertCurrency(amount, requestData.From, requestData.To)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := ConversionResponse{
		Amount: requestData.Amount,
		From:   requestData.From,
		To:     requestData.To,
		Result: result,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterRoutes(r *gin.Engine) {
	r.POST("/convert", ConvertHandler)
}
