package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Email          string `bson:"email,omitempty"`
	Password       string `bson:"password,omitempty"`
	Username       string `bson:"username,omitempty"`
	Validated      bool   `bson:"validated"`
	ValidationCode string `bson:"validation_code,omitempty"`
	ProfilePic     string `bson:"profile_pic,omitempty"`
	ID             string `bson:"_id,omitempty"`
}

func (u *User) Clone() *User {
	return &User{
		Email:          u.Email,
		Password:       u.Password,
		Username:       u.Username,
		Validated:      u.Validated,
		ValidationCode: u.ValidationCode,
		ProfilePic:     u.ProfilePic,
		ID:             u.ID,
	}
}

func (u *User) PurgedBson() bson.M {

	purgedBson := bson.M{}

	if u.Email != "" {
		purgedBson["email"] = u.Email
	}

	if u.Password != "" {
		purgedBson["password"] = u.Password
	}

	if u.Username != "" {
		purgedBson["username"] = u.Username
	}

	if u.Validated {
		purgedBson["validated"] = u.Validated
	}

	if u.ValidationCode != "" {
		purgedBson["validation_code"] = u.ValidationCode
	}

	if u.ProfilePic != "" {
		purgedBson["profile_pic"] = u.ProfilePic
	}

	if u.ID != "" {
		purgedBson["_id"] = u.ID
	}

	return purgedBson
}
