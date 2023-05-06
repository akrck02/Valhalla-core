package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Validated bool               `bson:"validated"`
	ID        primitive.ObjectID `bson:"_id,omitempty"`
}
