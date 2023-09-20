package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" validate:"required"`
	Username  string             `json:"username" bson:"username" validate:"required"`
	Lastname  string             `json:"lastname" bson:"lastName" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Password  string             `json:"password" bson:"password" validate:"required,min=8"`
	Avatar    []byte             `json:"avatar" bson:"avatar"`
	Roles     []string           `json:"roles" bson:"roles"`
	Address   string             `json:"address" bson:"address"`
	CreatedAt time.Time          `json:"created_at" bson:"createdAt"`
}
