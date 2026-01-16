package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/m-bromo/dungeon-desk/character-service/internal/database/sqlc"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain/entities"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain/repository"
)

func NewClassRepository(querier sqlc.Querier) repository.ClassRepository {
	return &classRepository{
		querier: querier,
	}
}

type classRepository struct {
	querier sqlc.Querier
}

func (r *classRepository) GetCharacterClasses(ctx context.Context, characterID uuid.UUID) ([]*entities.Class, error) {
	rows, err := r.querier.GetCharacterClasses(ctx, characterID)
	if err != nil {
		return nil, err
	}

	if len(rows) <= 0 {
		return nil, errors.New("No classes found for this character")
	}

	var classes []*entities.Class
	for _, row := range rows {
		classes = append(classes, &entities.Class{
			ID:          int(row.ID),
			Name:        row.Name,
			Description: row.Name,
			Level:       int(row.ClassLevel.Int32),
		})
	}

	return classes, nil
}
