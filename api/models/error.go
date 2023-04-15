package models

type Error struct {
	Code    int    `bson:"code,omitempty"`
	Error   int    `bson:"error,omitempty"`
	Message string `bson:"message,omitempty"`
}
