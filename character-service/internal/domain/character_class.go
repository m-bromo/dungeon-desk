package domain

import "github.com/google/uuid"

type CharacterClass struct {
	CharacterID uuid.UUID
	ClassID     int
	ClassLevel  int
}
