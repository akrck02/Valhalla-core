package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	Name        string               `bson:"name,omitempty"`
	Description string               `bson:"description,omitempty"`
	ProfilePic  string               `bson:"profilepic,omitempty"`
	Projects    []primitive.ObjectID `bson:"projects,omitempty"`
	Owner       primitive.ObjectID   `bson:"owner,omitempty"`
	Members     []primitive.ObjectID `bson:"members,omitempty"`
}
