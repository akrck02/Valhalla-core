package models

type Response struct {
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
