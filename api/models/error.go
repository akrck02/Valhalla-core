package models

type Error struct {
	Status  int    `json:"status,omitempty"`
	Error   int    `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
