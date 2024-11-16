package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CalculateInvestmentRatio(c *gin.Context) {
	log.Println("Starting investment ratio calculation")

	// Get form values and trim any whitespace
	totalInvestmentStr := strings.TrimSpace(c.PostForm("total_investment"))
	netWorthStr := strings.TrimSpace(c.PostForm("net_worth"))

	log.Printf("Received values - Total Investment: %s, Net Worth: %s\n", totalInvestmentStr, netWorthStr)

	// Check for empty inputs
	if totalInvestmentStr == "" || netWorthStr == "" {
		log.Println("Error: Empty input values")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "All fields are required",
		})
		return
	}

	// Parse total investment
	totalInvestment, err1 := strconv.ParseFloat(totalInvestmentStr, 64)
	if err1 != nil {
		log.Printf("Error parsing total investment: %v\n", err1)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  fmt.Sprintf("Invalid total investment value: %s", totalInvestmentStr),
		})
		return
	}

	// Parse net worth
	netWorth, err2 := strconv.ParseFloat(netWorthStr, 64)
	if err2 != nil {
		log.Printf("Error parsing net worth: %v\n", err2)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  fmt.Sprintf("Invalid net worth value: %s", netWorthStr),
		})
		return
	}

	log.Printf("Parsed values - Total Investment: %.2f, Net Worth: %.2f\n", totalInvestment, netWorth)

	// Validate values
	if netWorth == 0 {
		log.Println("Error: Net worth is zero")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Net worth must not be zero",
		})
		return
	}

	if totalInvestment < 0 {
		log.Println("Error: Total investment is negative")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Total investment cannot be negative",
		})
		return
	}

	// Calculate the ratio
	percentage := (totalInvestment / netWorth) * 100
	log.Printf("Calculated percentage: %.2f%%\n", percentage)

	// Determine recommendation based on general financial guidelines
	var recommendation string
	var message string

	if netWorth < 0 {
		recommendation = "red"
		message = fmt.Sprintf("Your investment ratio is %.1f%%. However, your net worth is negative. Focus on building positive net worth before increasing investments.", percentage)
	} else if percentage < 20 {
		recommendation = "yellow"
		message = fmt.Sprintf("Your investment ratio is %.1f%%. Consider increasing your investments to build long-term wealth.", percentage)
	} else if percentage < 40 {
		recommendation = "green"
		message = fmt.Sprintf("Your investment ratio is %.1f%%. Good job! You have a healthy portion of your net worth invested.", percentage)
	} else if percentage < 60 {
		recommendation = "green"
		message = fmt.Sprintf("Your investment ratio is %.1f%%. Excellent! You have a strong investment portfolio relative to your net worth.", percentage)
	} else {
		recommendation = "yellow"
		message = fmt.Sprintf("Your investment ratio is %.1f%%. While having substantial investments is good, consider maintaining some liquid assets for emergencies.", percentage)
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
