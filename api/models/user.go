package models

type User struct {
	Email     string `bson:"email,omitempty"`
	Password  string `bson:"password,omitempty"`
	Username  string `bson:"username,omitempty"`
	Validated bool   `bson:"validated"`
}
