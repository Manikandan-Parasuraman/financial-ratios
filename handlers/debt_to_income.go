package handlers

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
)

func CalculateDebtToIncome(c *gin.Context) {
    // Get form values and trim any whitespace
    totalMonthlyDebtStr := strings.TrimSpace(c.PostForm("total_monthly_debt"))
    grossIncomeStr := strings.TrimSpace(c.PostForm("gross_income"))

    fmt.Printf("Received values - Monthly Debt: '%s', Gross Income: '%s'\n", totalMonthlyDebtStr, grossIncomeStr)

    // Check for empty inputs
    if totalMonthlyDebtStr == "" || grossIncomeStr == "" {
        fmt.Println("Error: Empty input values")
        c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
        return
    }

    // Parse monthly debt
    monthlyDebt, err1 := strconv.ParseFloat(totalMonthlyDebtStr, 64)
    if err1 != nil {
        fmt.Printf("Error parsing monthly debt: %v\n", err1)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid monthly debt value: %s", totalMonthlyDebtStr)})
        return
    }

    // Parse gross income
    grossIncome, err2 := strconv.ParseFloat(grossIncomeStr, 64)
    if err2 != nil {
        fmt.Printf("Error parsing gross income: %v\n", err2)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid gross income value: %s", grossIncomeStr)})
        return
    }

    fmt.Printf("Parsed values - Monthly Debt: %.2f, Gross Income: %.2f\n", monthlyDebt, grossIncome)

    // Validate values
    if grossIncome <= 0 {
        fmt.Println("Error: Gross income is zero or negative")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Gross income must be greater than zero"})
        return
    }

    if monthlyDebt < 0 {
        fmt.Println("Error: Monthly debt is negative")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Monthly debt cannot be negative"})
        return
    }

    // Calculate the ratio
    percentage := (monthlyDebt / grossIncome) * 100
    fmt.Printf("Calculated percentage: %.2f%%\n", percentage)

    // Determine recommendation
    recommendation := "green"
    var message string

    if percentage > 36 {
        recommendation = "red"
        message = fmt.Sprintf("Your debt-to-income ratio is %.1f%%, which is high. Consider debt reduction strategies.", percentage)
    } else if percentage > 28 {
        recommendation = "yellow"
        message = fmt.Sprintf("Your debt-to-income ratio is %.1f%%. Watch your debt levels and try to reduce non-essential borrowing.", percentage)
    } else {
        message = fmt.Sprintf("Your debt-to-income ratio is %.1f%%, which is in a healthy range.", percentage)
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
