package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CalculateSavingsToIncome(c *gin.Context) {
    savingsStr := c.PostForm("savings")
    grossIncomeStr := c.PostForm("gross_income")

    savings, err1 := strconv.ParseFloat(savingsStr, 64)
    grossIncome, err2 := strconv.ParseFloat(grossIncomeStr, 64)

    if err1 != nil || err2 != nil || grossIncome <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    percentage := (savings / grossIncome) * 100
    recommendation := "green"
    if percentage < 20 {
        recommendation = "red"
    }

    c.JSON(http.StatusOK, gin.H{
        "percentage":     percentage,
        "recommendation": recommendation,
    })
}
