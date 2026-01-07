package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCharacterAtMaxLevel = errors.New("the character has achieved the max level")
)

const (
	MaxLevel = 20
)

type Character struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Level        int
	Experience   int
	ArmorClass   int
	HitPoints    int
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Classes      []*CharacterClass
	Traits       []*Trait
	Flaws        []*Flaw
}
