package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CalculateNetWorthToIncome(c *gin.Context) {
	log.Println("Starting net worth to income calculation")

	// Get form values and trim any whitespace
	netWorthStr := strings.TrimSpace(c.PostForm("net_worth"))
	annualIncomeStr := strings.TrimSpace(c.PostForm("annual_income"))

	log.Printf("Received values - Net Worth: %s, Annual Income: %s\n", netWorthStr, annualIncomeStr)

	// Check for empty inputs
	if netWorthStr == "" || annualIncomeStr == "" {
		log.Println("Error: Empty input values")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "All fields are required",
		})
		return
	}

	// Parse net worth
	netWorth, err1 := strconv.ParseFloat(netWorthStr, 64)
	if err1 != nil {
		log.Printf("Error parsing net worth: %v\n", err1)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  fmt.Sprintf("Invalid net worth value: %s", netWorthStr),
		})
		return
	}

	// Parse annual income
	annualIncome, err2 := strconv.ParseFloat(annualIncomeStr, 64)
	if err2 != nil {
		log.Printf("Error parsing annual income: %v\n", err2)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  fmt.Sprintf("Invalid annual income value: %s", annualIncomeStr),
		})
		return
	}

	log.Printf("Parsed values - Net Worth: %.2f, Annual Income: %.2f\n", netWorth, annualIncome)

	// Validate values
	if annualIncome <= 0 {
		log.Println("Error: Annual income is zero or negative")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Annual income must be greater than zero",
		})
		return
	}

	// Calculate the ratio
	percentage := (netWorth / annualIncome) * 100
	log.Printf("Calculated percentage: %.2f%%\n", percentage)

	// Determine recommendation based on general financial guidelines
	var recommendation string
	var message string

	if percentage < 0 {
		recommendation = "red"
		message = fmt.Sprintf("Your net worth to income ratio is %.1f%%. A negative ratio indicates more liabilities than assets. Consider debt reduction strategies and building savings.", percentage)
	} else if percentage < 100 {
		recommendation = "yellow"
		message = fmt.Sprintf("Your net worth to income ratio is %.1f%%. While positive, aim to build your net worth to at least 1x your annual income.", percentage)
	} else if percentage < 300 {
		recommendation = "green"
		message = fmt.Sprintf("Your net worth to income ratio is %.1f%%. Good job! Your net worth is greater than your annual income.", percentage)
	} else {
		recommendation = "green"
		message = fmt.Sprintf("Your net worth to income ratio is %.1f%%. Excellent! Your net worth is significantly higher than your annual income, indicating strong financial health.", percentage)
	}

	response := gin.H{
		"status":         "success",
		"percentage":     percentage,
		"recommendation": recommendation,
		"message":        message,
	}

	log.Printf("Sending response: %+v\n", response)
	c.JSON(http.StatusOK, response)
}
