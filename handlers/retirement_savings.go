package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CalculateRetirementSavings(c *gin.Context) {
	// Get form values and trim any whitespace
	retirementSavingsStr := strings.TrimSpace(c.PostForm("retirement_savings"))
	annualIncomeStr := strings.TrimSpace(c.PostForm("annual_income_retirement"))

	fmt.Printf("Received values - Retirement Savings: '%s', Annual Income: '%s'\n", retirementSavingsStr, annualIncomeStr)

	// Check for empty inputs
	if retirementSavingsStr == "" || annualIncomeStr == "" {
		fmt.Println("Error: Empty input values")
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Parse retirement savings
	retirementSavings, err1 := strconv.ParseFloat(retirementSavingsStr, 64)
	if err1 != nil {
		fmt.Printf("Error parsing retirement savings: %v\n", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid retirement savings value: %s", retirementSavingsStr)})
		return
	}

	// Parse annual income
	annualIncome, err2 := strconv.ParseFloat(annualIncomeStr, 64)
	if err2 != nil {
		fmt.Printf("Error parsing annual income: %v\n", err2)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid annual income value: %s", annualIncomeStr)})
		return
	}

	fmt.Printf("Parsed values - Retirement Savings: %.2f, Annual Income: %.2f\n", retirementSavings, annualIncome)

	// Validate values
	if annualIncome <= 0 {
		fmt.Println("Error: Annual income is zero or negative")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Annual income must be greater than zero"})
		return
	}

	if retirementSavings < 0 {
		fmt.Println("Error: Retirement savings is negative")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Retirement savings cannot be negative"})
		return
	}

	// Calculate the ratio
	percentage := (retirementSavings / annualIncome) * 100
	fmt.Printf("Calculated percentage: %.2f%%\n", percentage)

	// Determine recommendation
	recommendation := "green"
	var message string

	if percentage < 5 {
		recommendation = "red"
		message = fmt.Sprintf("Your retirement savings ratio is %.1f%%, which is low. Consider increasing your retirement contributions.", percentage)
	} else if percentage < 10 {
		recommendation = "yellow"
		message = fmt.Sprintf("Your retirement savings ratio is %.1f%%. Try to increase contributions to at least 10%% of your income.", percentage)
	} else {
		message = fmt.Sprintf("Your retirement savings ratio is %.1f%%, which is in a healthy range.", percentage)
	}

	// Prepare and send response
	response := gin.H{
		"percentage":     percentage,
		"recommendation": recommendation,
		"message":        message,
	}
	fmt.Printf("Sending response: %+v\n", response)
	c.JSON(http.StatusOK, response)
}
