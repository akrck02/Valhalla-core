package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	ProfilePic  string             `bson:"profilepic,omitempty"`
	Projects    []string           `bson:"projects,omitempty"`
	Owner       string             `bson:"owner,omitempty"`
	Members     []string           `bson:"members,omitempty"`
	ID          primitive.ObjectID `bson:"_id,omitempty"`
}
