package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RunGinImpl() {
	r := gin.Default()

	r.POST("/receipts/process", ginProcessReceipt)
	r.GET("/receipts/:id/points", ginGetPoints)

	r.Run(":8080")
}

func ginProcessReceipt(c *gin.Context) {
	var receipt Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uuid.NewString()
	points := calculatePoints(receipt)
	db[id] = points
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func ginGetPoints(c *gin.Context) {
	id := c.Param("id")
	if points, exists := db[id]; exists {
		c.JSON(http.StatusOK, gin.H{"points": points})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
	}
}
