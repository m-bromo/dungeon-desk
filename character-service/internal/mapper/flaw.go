package mapper

import (
	"github.com/google/uuid"
	db "github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain"
)

func ToCreateFlawParams(d *domain.Flaw, characterID uuid.UUID) db.CreateFlawParams {
	return db.CreateFlawParams{
		ID:          d.ID,
		CharacterID: characterID,
		Description: d.Description,
	}
}
