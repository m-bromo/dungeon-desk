package dto

type CreateCharacterPayload struct {
	Name         string                        `json:"name" binding:"required"`
	Description  string                        `json:"description"`
	Level        int                           `json:"level" binding:"required,min=1"`
	ArmorClass   int                           `json:"armorClass" binding:"required"`
	HitPoints    int                           `json:"hitPoints" binding:"required"`
	Strength     int                           `json:"strength" binding:"required"`
	Dexterity    int                           `json:"dexterity" binding:"required"`
	Constitution int                           `json:"constitution" binding:"required"`
	Intelligence int                           `json:"intelligence" binding:"required"`
	Wisdom       int                           `json:"wisdom" binding:"required"`
	Charisma     int                           `json:"charisma" binding:"required"`
	Classes      []CreateCharacterClassPayload `json:"classes" binding:"required,dive"`
	Traits       []string                      `json:"traits"`
	Flaws        []string                      `json:"flaws"`
}
