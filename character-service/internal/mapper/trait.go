package mapper

import (
	"github.com/google/uuid"
	db "github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain"
)

func ToCreateTraitParams(d *domain.Trait, characterID uuid.UUID) db.CreateTraitParams {
	return db.CreateTraitParams{
		ID:          d.ID,
		CharacterID: characterID,
		Description: d.Description,
	}
}
