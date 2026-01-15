package models

import "github.com/m-bromo/dungeon-desk/character-service/internal/domain"

type CreateCharacterPayload struct {
	Name         string           `json:"name" binding:"required"`
	Description  string           `json:"description"`
	Alignment    domain.Alignment `json:"alignment"`
	ClassID      int              `json:"class_id"`
	Level        int              `json:"level"`
	Strength     int              `json:"strength" binding:"required"`
	Dexterity    int              `json:"dexterity" binding:"required"`
	Constitution int              `json:"constitution" binding:"required"`
	Intelligence int              `json:"intelligence" binding:"required"`
	Wisdom       int              `json:"wisdom" binding:"required"`
	Charisma     int              `json:"charisma" binding:"required"`
	Traits       []string         `json:"traits"`
	Flaws        []string         `json:"flaws"`
}
