package utils

import "go.mongodb.org/mongo-driver/mongo"

// Database represent the mongo connection
type Database struct {
	Client *mongo.Client
}
