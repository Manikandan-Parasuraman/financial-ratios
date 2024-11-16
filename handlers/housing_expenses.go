package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CalculateHousingExpenses(c *gin.Context) {
	// Get form values and trim any whitespace
	monthlyHousingCostStr := strings.TrimSpace(c.PostForm("monthly_housing_cost"))
	monthlyGrossIncomeStr := strings.TrimSpace(c.PostForm("monthly_gross_income"))

	fmt.Printf("Received values - Housing Cost: '%s', Monthly Gross Income: '%s'\n", monthlyHousingCostStr, monthlyGrossIncomeStr)

	// Check for empty inputs
	if monthlyHousingCostStr == "" || monthlyGrossIncomeStr == "" {
		fmt.Println("Error: Empty input values")
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Parse housing cost
	housingCost, err1 := strconv.ParseFloat(monthlyHousingCostStr, 64)
	if err1 != nil {
		fmt.Printf("Error parsing housing cost: %v\n", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid housing cost value: %s", monthlyHousingCostStr)})
		return
	}

	// Parse monthly gross income
	grossIncome, err2 := strconv.ParseFloat(monthlyGrossIncomeStr, 64)
	if err2 != nil {
		fmt.Printf("Error parsing monthly gross income: %v\n", err2)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid monthly gross income value: %s", monthlyGrossIncomeStr)})
		return
	}

	fmt.Printf("Parsed values - Housing Cost: %.2f, Monthly Gross Income: %.2f\n", housingCost, grossIncome)

	// Validate values
	if grossIncome <= 0 {
		fmt.Println("Error: Monthly gross income is zero or negative")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Monthly gross income must be greater than zero"})
		return
	}

	if housingCost < 0 {
		fmt.Println("Error: Housing cost is negative")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Housing cost cannot be negative"})
		return
	}

	// Calculate the ratio
	percentage := (housingCost / grossIncome) * 100
	fmt.Printf("Calculated percentage: %.2f%%\n", percentage)

	// Determine recommendation based on the 28% rule
	recommendation := "green"
	var message string

	if percentage > 35 {
		recommendation = "red"
		message = fmt.Sprintf("Your housing expenses ratio is %.1f%%, which is significantly high. Consider ways to reduce housing costs or increase income.", percentage)
	} else if percentage > 28 {
		recommendation = "yellow"
		message = fmt.Sprintf("Your housing expenses ratio is %.1f%%. It's slightly above the recommended 28%%. Consider budgeting to reduce housing costs.", percentage)
	} else {
		message = fmt.Sprintf("Your housing expenses ratio is %.1f%%, which is within the recommended range (below 28%%).", percentage)
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
