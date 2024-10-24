package main

import (
	"context"
	"log"
	"net/http"
	"request-broker/internal/handlers"
	"request-broker/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.TODO())

	db := client.Database("request-broker").Collection("queue")
	archive := client.Database("request-broker").Collection("archive")

	go services.ProcessQueue(db, archive)

	r := gin.Default()
	r.POST("/queue", func(c *gin.Context) {
		handlers.AddToQuequeHandler(c, db)
	})
	r.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Service is running"})
	})

	r.Run(":8080")
	log.Println("Server request-broker running on port 8080")
}
