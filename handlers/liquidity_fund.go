package handlers

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"
    "log"

    "github.com/gin-gonic/gin"
)

func CalculateLiquidityFund(c *gin.Context) {
    // Log the entire request details
    log.Printf("Received request: %+v", c.Request)
    log.Printf("Request Headers: %+v", c.Request.Header)

    // Get form values and trim any whitespace
    liquidityFundStr := strings.TrimSpace(c.PostForm("liquidity_fund"))
    netWorthStr := strings.TrimSpace(c.PostForm("net_worth_liquidity"))

    log.Printf("Received form values - Liquidity Fund: '%s', Net Worth: '%s'", liquidityFundStr, netWorthStr)

    // Check for empty inputs
    if liquidityFundStr == "" || netWorthStr == "" {
        log.Println("Error: Empty input values")
        c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
        return
    }

    // Parse liquidity fund
    liquidityFund, err1 := strconv.ParseFloat(liquidityFundStr, 64)
    if err1 != nil {
        log.Printf("Error parsing liquidity fund: %v", err1)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid liquidity fund value: %s", liquidityFundStr)})
        return
    }

    // Parse net worth
    netWorth, err2 := strconv.ParseFloat(netWorthStr, 64)
    if err2 != nil {
        log.Printf("Error parsing net worth: %v", err2)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid net worth value: %s", netWorthStr)})
        return
    }

    log.Printf("Parsed values - Liquidity Fund: %.2f, Net Worth: %.2f", liquidityFund, netWorth)

    // Validate values
    if netWorth <= 0 {
        log.Println("Error: Net worth is zero or negative")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Net worth must be greater than zero"})
        return
    }

    if liquidityFund < 0 {
        log.Println("Error: Liquidity fund is negative")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Liquidity fund cannot be negative"})
        return
    }

    // Calculate the ratio
    percentage := (liquidityFund / netWorth) * 100
    log.Printf("Calculated percentage: %.2f%%", percentage)

    // Determine recommendation
    recommendation := "green"
    var message string

    if percentage < 10 {
        recommendation = "red"
        message = fmt.Sprintf("Your liquidity fund ratio is %.1f%%, which is low. Consider building up more liquid assets for emergencies.", percentage)
    } else if percentage < 20 {
        recommendation = "yellow"
        message = fmt.Sprintf("Your liquidity fund ratio is %.1f%%. Consider increasing your liquid assets slightly.", percentage)
    } else {
        message = fmt.Sprintf("Your liquidity fund ratio is %.1f%%, which is in a healthy range.", percentage)
    }

    // Prepare and send response
    response := gin.H{
        "percentage":     percentage,
        "recommendation": recommendation,
        "message":        message,
    }
    
    log.Printf("Sending response: %+v", response)
    c.JSON(http.StatusOK, response)
}
