package models

type Device struct {
	User      string `bson:"user,omitempty"`
	Address   string `bson:"address,omitempty"`
	UserAgent string `bson:"useragent,omitempty"`
	Token     string `bson:"token,omitempty"`
}
