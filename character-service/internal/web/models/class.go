package models

type ClassPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}
