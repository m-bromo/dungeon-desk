package domain

import "github.com/google/uuid"

type Flaw struct {
	ID          uuid.UUID
	Description string
}
