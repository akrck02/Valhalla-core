package models

type Request struct {
	Authorization string `json:"authorization"`
	IP            string `json:"ip"`
	UserAgent     string `json:"userAgent"`
	User          *User  `json:"user"`
}
