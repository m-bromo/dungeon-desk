package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain"
)

type CharacterRepository interface {
	CreateCharacter(ctx context.Context, character *domain.Character) error
	AddClassToCharacter(ctx context.Context, characterID uuid.UUID, classID int) error
}

type characterRepository struct {
	querier sqlc.Querier
}

func NewCharacterRepository(querier sqlc.Querier) CharacterRepository {
	return &characterRepository{
		querier: querier,
	}
}

func (r *characterRepository) CreateCharacter(ctx context.Context, character *domain.Character) error {
	_, err := r.querier.CreateCharacter(ctx, sqlc.CreateCharacterParams{
		ID:   character.ID,
		Name: character.Name,
		Description: sql.NullString{
			String: character.Description,
			Valid:  true,
		},
		Alignment: sqlc.NullAlignment{
			Alignment: sqlc.Alignment(character.Alignment),
			Valid:     true,
		},
		TotalLevel:       int32(character.TotalLevel),
		Experience:       int32(character.Experience),
		ArmorClass:       int32(character.ArmorClass),
		HitPoints:        int32(character.HitPoints),
		CurrentHitPoints: int32(character.CurrentHitPoints),
		Strength:         int32(character.Strength),
		Dexterity:        int32(character.Dexterity),
		Constitution:     int32(character.Constitution),
		Intelligence:     int32(character.Intelligence),
		Wisdom:           int32(character.Wisdom),
		Charisma:         int32(character.Charisma),
		Traits:           character.Traits,
		Flaws:            character.Flaws,
	})

	return err
}
func (r *characterRepository) AddClassToCharacter(ctx context.Context, characterID uuid.UUID, classID int) error {
	_, err := r.querier.AddClassToCharacter(ctx, sqlc.AddClassToCharacterParams{
		CharacterID: characterID,
		ClassID:     int32(classID),
		ClassLevel: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
	})

	return err
}
