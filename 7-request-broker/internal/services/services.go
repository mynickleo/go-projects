package services

import (
	"context"
	"log"
	"net/http"
	"request-broker/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProcessQueue(db *mongo.Collection, archive *mongo.Collection) {
	for {
		var req models.Request
		err := db.FindOne(context.Background(), bson.M{"status": "pending"}).Decode(&req)
		if err != nil {
			log.Println("No pending requests")
			time.Sleep(10 * time.Second)
			continue
		}

		httpReq, err := http.NewRequest(req.Method, req.URL, nil)
		if err != nil {
			log.Println("Error creating HTTP request:", err)
			logError(req, db)
			continue
		}

		client := &http.Client{}
		resp, err := client.Do(httpReq)
		if err != nil || resp.StatusCode >= 400 {
			log.Println("Error sending HTTP request or bad status code:", err)
			logError(req, db)
			continue
		}

		archiveRequest(req, "completed", db, archive)
	}
}

func logError(req models.Request, db *mongo.Collection) {
	req.RetryCount++
	req.UpdatedAt = time.Now()

	if req.RetryCount >= 5 {
		archiveRequest(req, "error", db, nil)
	} else {
		_, err := db.UpdateOne(context.Background(), bson.M{"_id": req.ID}, bson.M{
			"$set": bson.M{
				"retry_count": req.RetryCount,
				"updated_at":  req.UpdatedAt,
			},
		})
		if err != nil {
			log.Println("Failed to update request retry count:", err)
		}
	}
}

func archiveRequest(req models.Request, status string, db *mongo.Collection, archive *mongo.Collection) {
	req.Status = status
	req.UpdatedAt = time.Now()

	if archive != nil {
		_, err := archive.InsertOne(context.Background(), req)
		if err != nil {
			log.Println("Failed to archive request:", err)
		}
	}

	_, err := db.DeleteOne(context.Background(), bson.M{"_id": req.ID})
	if err != nil {
		log.Println("Failed to delete request from queue:", err)
	}
}
