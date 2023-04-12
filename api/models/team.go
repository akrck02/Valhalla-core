package models

type Team struct {
	Name        string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	ProfilePic  string `bson:"profilepic,omitempty"`
}
