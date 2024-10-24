package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	URL        string             `bson:"url"`
	Method     string             `bson:"method"`
	Body       interface{}        `bson:"body,omitempty"`
	Query      map[string]string  `bson:"query,omitempty"`
	Status     string             `bson:"status"`
	RetryCount int                `bson:"retry_count"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}
