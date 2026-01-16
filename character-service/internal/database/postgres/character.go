package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain/entities"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain/repository"
)

type characterRepository struct {
	querier sqlc.Querier
}

func NewCharacterRepository(querier sqlc.Querier) repository.CharacterRepository {
	return &characterRepository{
		querier: querier,
	}
}

func (r *characterRepository) CreateCharacter(ctx context.Context, character *entities.Character) error {
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

func (r characterRepository) GetCharacter(ctx context.Context, characterID uuid.UUID) (*entities.Character, error) {
	character, err := r.querier.GetCharacter(ctx, characterID)
	if err != nil {
		return nil, err
	}

	return &entities.Character{
		Name:              character.Name,
		Description:       character.Description.String,
		Alignment:         entities.Alignment(character.Alignment.Alignment),
		TotalLevel:        int(character.TotalLevel),
		Experience:        int(character.Experience),
		ArmorClass:        int(character.ArmorClass),
		HitPoints:         int(character.HitPoints),
		CurrentHitPoints:  int(character.CurrentHitPoints),
		Speed:             int(character.Speed.Int32),
		ProeficiencyBonus: int(character.ProficiencyBonus.Int32),
		Strength:          int(character.Strength),
		Dexterity:         int(character.Dexterity),
		Constitution:      int(character.Constitution),
		Wisdom:            int(character.Wisdom),
		Intelligence:      int(character.Intelligence),
		Charisma:          int(character.Charisma),
		Traits:            character.Traits,
		Flaws:             character.Flaws,
	}, nil
}
