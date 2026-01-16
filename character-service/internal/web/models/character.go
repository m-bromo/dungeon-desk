package models

import domain "github.com/m-bromo/dungeon-desk/character-service/internal/domain/entities"

type CreateCharacterPayload struct {
	Name         string           `json:"name" binding:"required"`
	Description  string           `json:"description"`
	Alignment    domain.Alignment `json:"alignment"`
	ClassID      int              `json:"class_id"`
	TotalLevel   int              `json:"total_level"`
	Strength     int              `json:"strength" binding:"required"`
	Dexterity    int              `json:"dexterity" binding:"required"`
	Constitution int              `json:"constitution" binding:"required"`
	Intelligence int              `json:"intelligence" binding:"required"`
	Wisdom       int              `json:"wisdom" binding:"required"`
	Charisma     int              `json:"charisma" binding:"required"`
	Traits       []string         `json:"traits"`
	Flaws        []string         `json:"flaws"`
}

type GetDetailedCharacterResponse struct {
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Alignment    domain.Alignment `json:"alignment"`
	Classes      []ClassPayload   `json:"classes"`
	TotalLevel   int              `json:"total_level"`
	Strength     int              `json:"strength"`
	Dexterity    int              `json:"dexterity"`
	Constitution int              `json:"constitution"`
	Intelligence int              `json:"intelligence"`
	Wisdom       int              `json:"wisdom"`
	Charisma     int              `json:"charisma"`
	Traits       []string         `json:"traits"`
	Flaws        []string         `json:"flaws"`
}
