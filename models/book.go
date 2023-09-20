package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Author      string             `json:"author" bson:"author"`
	Published   time.Time          `json:"published" bson:"published"`
	Genre       []string           `json:"genre" bson:"genre"`
	CreatedAt   time.Time          `json:"created_at" bson:"createdAt"`
}
