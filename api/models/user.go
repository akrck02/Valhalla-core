package models

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
