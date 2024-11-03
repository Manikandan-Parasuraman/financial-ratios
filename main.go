package main

import (
	"go/financial-ratios/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    router.Static("/static", "./static")

    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })
    router.POST("/calculate-emergency-fund", handlers.CalculateEmergencyFund)
    router.POST("/calculate-savings-to-income", handlers.CalculateSavingsToIncome)
    router.POST("/calculate-debt-to-income", handlers.CalculateDebtToIncome)

    router.Run(":8080")
}
