package mapper

import (
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain/entities"
	"github.com/m-bromo/dungeon-desk/character-service/internal/web/models"
)

func ToDomainInput(payload *models.CreateCharacterPayload) *entities.CreateCharacterInput {
	return &entities.CreateCharacterInput{
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

func ToCharacterResponse(output *entities.GetCharacterOutput) *models.GetDetailedCharacterResponse {
	var classPayload []models.ClassPayload
	for _, class := range output.Classes {
		classPayload = append(classPayload, models.ClassPayload{
			Name:        class.Name,
			Description: class.Description,
			Level:       class.Level,
		})
	}

	return &models.GetDetailedCharacterResponse{
		Name:         output.Name,
		Description:  output.Description,
		Alignment:    output.Alignment,
		Classes:      classPayload,
		TotalLevel:   output.TotalLevel,
		Strength:     output.Strength,
		Dexterity:    output.Dexterity,
		Constitution: output.Constitution,
		Intelligence: output.Intelligence,
		Wisdom:       output.Wisdom,
		Charisma:     output.Charisma,
		Traits:       output.Traits,
		Flaws:        output.Flaws,
	}
}
