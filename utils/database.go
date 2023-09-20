package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// this function create a new instance of database.
func NewDatabase() (*Database, error) {
	mongoURI := os.Getenv("MONGODB_URI")

	// set up options
	clientOptions := options.Client().ApplyURI(mongoURI)
	clientOptions.SetMaxPoolSize(10) // Max number connection

	// establish connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection to MongoDB is established correctly!")

	return &Database{Client: client}, nil
}

// close the mongo connection
func (db *Database) Close() {
	if db.Client != nil {
		db.Client.Disconnect(context.Background())
	}
}

func (db *Database) GetName() string {
	return os.Getenv("DATABASE_NAME")
}
