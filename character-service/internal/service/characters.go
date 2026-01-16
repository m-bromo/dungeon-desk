package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain/entities"
	"github.com/m-bromo/dungeon-desk/character-service/internal/domain/repository"
)

type CharacterService interface {
	CreateCharacter(ctx context.Context, input *entities.CreateCharacterInput) error
	GetCharacter(ctx context.Context, characterID uuid.UUID) (*entities.GetCharacterOutput, error)
}

type characterService struct {
	characterRepository repository.CharacterRepository
	classRepository     repository.ClassRepository
}

func NewCharacterService(
	characterRepository repository.CharacterRepository,
	classRepository repository.ClassRepository,
) CharacterService {
	return &characterService{
		characterRepository: characterRepository,
		classRepository:     classRepository,
	}
}

func (s *characterService) CreateCharacter(ctx context.Context, input *entities.CreateCharacterInput) error {
	character := entities.NewCharacter(input)

	if err := s.characterRepository.CreateCharacter(ctx, character); err != nil {
		slog.Error("Failed to create character", "error", err.Error())
		return err
	}

	if err := s.characterRepository.AddClassToCharacter(ctx, character.ID, input.ClassID); err != nil {
		slog.Error("Failed to add class to character", "error", err.Error())
		return err
	}

	return nil
}

func (s *characterService) GetCharacter(ctx context.Context, characterID uuid.UUID) (*entities.GetCharacterOutput, error) {
	character, err := s.characterRepository.GetCharacter(ctx, characterID)
	if err != nil {
		slog.Error("Failed to find character", "error", err.Error())
		return nil, err
	}

	character.Classes, err = s.classRepository.GetCharacterClasses(ctx, characterID)
	if err != nil {
		slog.Error("Failed to find character classes", "error", err.Error())
	}

	return &entities.GetCharacterOutput{
		Name:              character.Name,
		Description:       character.Description,
		Alignment:         character.Alignment,
		TotalLevel:        character.TotalLevel,
		Experience:        character.Experience,
		ProeficiencyBonus: character.ProeficiencyBonus,
		ArmorClass:        character.ArmorClass,
		HitPoints:         character.HitPoints,
		CurrentHitPoints:  character.CurrentHitPoints,
		Strength:          character.Strength,
		Dexterity:         character.Dexterity,
		Constitution:      character.Constitution,
		Wisdom:            character.Wisdom,
		Intelligence:      character.Intelligence,
		Charisma:          character.Charisma,
		Classes:           character.Classes,
	}, nil
}
