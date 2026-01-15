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
	ID                uuid.UUID
	Name              string
	Description       string
	Alignment         Alignment
	TotalLevel        int
	Experience        int
	ArmorClass        int
	HitPoints         int
	CurrentHitPoints  int
	Speed             int
	ProeficiencyBonus int
	Strength          int
	Dexterity         int
	Constitution      int
	Intelligence      int
	Wisdom            int
	Charisma          int
	Traits            []string
	Flaws             []string
	Feats             []*Feat
	Spells            []*Spell
	Classes           []*Class
}

type CharacterInput struct {
	Name         string
	Description  string
	Alignment    Alignment
	ClassID      int
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	traits       []string
	flaws        []string
}

func NewCharacter(input *CharacterInput) *Character {
	baseArmorClass := 10 + (input.Dexterity - 10/2)
	baseHitPoints := 10 + (input.Constitution - 10/2)

	return &Character{
		ID:                uuid.New(),
		Name:              input.Name,
		Description:       input.Description,
		Alignment:         input.Alignment,
		TotalLevel:        1,
		Experience:        0,
		ProeficiencyBonus: 2,
		ArmorClass:        baseArmorClass,
		HitPoints:         baseHitPoints,
		CurrentHitPoints:  baseHitPoints,
		Strength:          input.Strength,
		Dexterity:         input.Dexterity,
		Constitution:      input.Constitution,
		Intelligence:      input.Intelligence,
		Wisdom:            input.Wisdom,
		Charisma:          input.Charisma,
		Traits:            input.traits,
		Flaws:             input.flaws,
	}
}
