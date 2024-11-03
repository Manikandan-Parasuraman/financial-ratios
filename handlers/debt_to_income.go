package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CalculateDebtToIncome(c *gin.Context) {
    totalMonthlyDebtStr := c.PostForm("total_debt")
    grossIncomeStr := c.PostForm("gross_income")

    monthlyDept, err1 := strconv.ParseFloat(totalMonthlyDebtStr, 64)
    grossIncome, err2 := strconv.ParseFloat(grossIncomeStr, 64)

    if err1 != nil || err2 != nil || monthlyDept <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    percentage := (monthlyDept / grossIncome) * 100
    fmt.Println("percentage :", percentage)

    recommendation := "green"
    if percentage < 15 {
        recommendation = "red"
    }

    c.JSON(http.StatusOK, gin.H{
        "percentage":     percentage,
        "recommendation": recommendation,
    })
}
