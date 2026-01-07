package mapper

import (
	"database/sql"

	"github.com/google/uuid"
	db "github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain"
	"github.com/m-bromo/dungeon-desk/character-service/internal/web/dto"
)

func ToCreateCharacterParams(d *domain.Character) db.CreateCharacterParams {
	return db.CreateCharacterParams{
		ID:   d.ID,
		Name: d.Name,
		Description: sql.NullString{
			String: d.Description,
			Valid:  true,
		},
		Level:        int32(d.Level),
		Experience:   int32(d.Experience),
		ArmorClass:   int32(d.ArmorClass),
		HitPoints:    int32(d.HitPoints),
		Strength:     int32(d.Strength),
		Dexterity:    int32(d.Dexterity),
		Constitution: int32(d.Constitution),
		Intelligence: int32(d.Intelligence),
		Wisdom:       int32(d.Wisdom),
		Charisma:     int32(d.Charisma),
	}
}

func ToCharacterDomain(payload dto.CreateCharacterPayload) *domain.Character {
	var classes []domain.CharacterClass
	for _, c := range payload.Classes {
		classes = append(classes, domain.CharacterClass{
			ClassID:    c.ClassID,
			ClassLevel: c.Level,
		})
	}

	var traits []domain.Trait
	for _, t := range payload.Traits {
		traits = append(traits, domain.Trait{
			ID:          uuid.New(),
			Description: t,
		})
	}

	var flaws []domain.Flaw
	for _, f := range payload.Flaws {
		flaws = append(flaws, domain.Flaw{
			ID:          uuid.New(),
			Description: f,
		})
	}

	return &domain.Character{
		ID:           uuid.New(),
		Name:         payload.Name,
		Description:  payload.Description,
		Level:        payload.Level,
		Experience:   payload.Experience,
		ArmorClass:   payload.ArmorClass,
		HitPoints:    payload.HitPoints,
		Strength:     payload.Strength,
		Dexterity:    payload.Dexterity,
		Constitution: payload.Constitution,
		Intelligence: payload.Intelligence,
		Wisdom:       payload.Wisdom,
		Charisma:     payload.Charisma,
		Classes:      classes,
		Traits:       traits,
		Flaws:        flaws,
	}
}
