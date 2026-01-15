package mapper

import (
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain"
	"github.com/m-bromo/dungeon-desk/character-service/internal/web/models"
)

func ToDomainInput(payload *models.CreateCharacterPayload) *domain.CreateCharacterInput {
	return &domain.CreateCharacterInput{
		Name:         payload.Name,
		Description:  payload.Description,
		Alignment:    payload.Alignment,
		ClassID:      payload.ClassID,
		Strength:     payload.Strength,
		Dexterity:    payload.Dexterity,
		Constitution: payload.Constitution,
		Wisdom:       payload.Wisdom,
		Intelligence: payload.Intelligence,
		Charisma:     payload.Charisma,
	}
}
