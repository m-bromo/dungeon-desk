package domain

import "github.com/google/uuid"

type Class struct {
	ID          uuid.UUID
	Level       int
	Name        string
	Description string
}
