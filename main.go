package main

import (
	"financial-ratios/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Enable debug mode
	gin.SetMode(gin.DebugMode)

	log.Println("Starting server initialization...")

	router := gin.Default()

	// Add middleware to log all requests
	router.Use(func(c *gin.Context) {
		log.Printf("Incoming %s request to %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
		log.Printf("Completed %s request to %s with status %d", c.Request.Method, c.Request.URL.Path, c.Writer.Status())
	})

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	router.GET("/calculator", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Add a test endpoint
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Register all calculator endpoints
	log.Println("Registering calculator endpoints...")

	router.POST("/calculator/calculate-emergency-fund", handlers.CalculateEmergencyFund)
	log.Println("Registered /calculator/calculate-emergency-fund")

	router.POST("/calculator/calculate-savings-to-income", handlers.CalculateSavingsToIncome)
	log.Println("Registered /calculator/calculate-savings-to-income")

	router.POST("/calculator/calculate-debt-to-income", handlers.CalculateDebtToIncome)
	log.Println("Registered /calculator/calculate-debt-to-income")

	router.POST("/calculator/calculate-housing-expenses", handlers.CalculateHousingExpenses)
	log.Println("Registered /calculator/calculate-housing-expenses")

	router.POST("/calculator/calculate-net-worth-income", handlers.CalculateNetWorthToIncome)
	log.Println("Registered /calculator/calculate-net-worth-income")

	router.POST("/calculator/calculate-investment-ratio", handlers.CalculateInvestmentRatio)
	log.Println("Registered /calculator/calculate-investment-ratio")

	router.POST("/calculator/calculate-retirement-savings", handlers.CalculateRetirementSavings)
	log.Println("Registered /calculator/calculate-retirement-savings")

	router.POST("/calculator/calculate-liquidity-fund", handlers.CalculateLiquidityFund)
	log.Println("Registered /calculator/calculate-liquidity-fund")

	// Add NoRoute handler for debugging
	router.NoRoute(func(c *gin.Context) {
		log.Printf("404 Not Found: %s %s", c.Request.Method, c.Request.URL.Path)
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Route not found",
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		})
	})

	// Print all registered routes
	log.Println("\nRegistered Routes:")
	for _, route := range router.Routes() {
		log.Printf("%s %s", route.Method, route.Path)
	}

	// Use port 8081 instead
	port := ":8081"
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
