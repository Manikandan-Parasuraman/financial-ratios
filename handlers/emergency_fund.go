package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CalculateEmergencyFund(c *gin.Context) {
    emergencyFundStr := c.PostForm("emergency_fund")
    monthlyExpensesStr := c.PostForm("monthly_expenses")

    emergencyFund, err1 := strconv.ParseFloat(emergencyFundStr, 64)
    monthlyExpenses, err2 := strconv.ParseFloat(monthlyExpensesStr, 64)

    if err1 != nil || err2 != nil || monthlyExpenses <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    months := int(emergencyFund / monthlyExpenses)
    recommendation := "green"
    if months < 3 {
        recommendation = "red"
    }

    c.JSON(http.StatusOK, gin.H{
        "months":        months,
        "recommendation": recommendation,
    })
}
