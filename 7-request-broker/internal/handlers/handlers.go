package handlers

import (
	"context"
	"net/http"
	"request-broker/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddToQuequeHandler(c *gin.Context, db *mongo.Collection) {
	var req models.Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req.Status = "pending"
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	req.RetryCount = 0

	_, err := db.InsertOne(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to queue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request added to queue"})
}
