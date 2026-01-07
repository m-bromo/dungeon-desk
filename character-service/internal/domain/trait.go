package domain

import "github.com/google/uuid"

type Trait struct {
	ID          uuid.UUID
	Description string
}
