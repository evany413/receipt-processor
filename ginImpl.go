package main

import (
	"github.com/gin-gonic/gin"
)

func RunGinImpl() {
	r := gin.Default()

	r.POST("/receipts/process", ginProcessReceipt)
	r.GET("/receipts/:id/points", ginGetPoints)

	r.Run(":8080")
}

func ginProcessReceipt(c *gin.Context) {
	// Implement your logic here
}

func ginGetPoints(c *gin.Context) {
	// Implement your logic here
}
