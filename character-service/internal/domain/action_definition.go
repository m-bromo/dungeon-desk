package domain

import "github.com/google/uuid"

type CostType string
type AttackType string

const (
	Action      CostType = "action"
	BonusAction CostType = "bonus_action"
	Reaction    CostType = "reaction"
	Movement    CostType = "movement"

	MeleeWeapon  AttackType = "melee_weapon"
	RangedWeapon AttackType = "ranged_weapon"
	MeleeSpell   AttackType = "melee_spell"
	RangedSpell  AttackType = "ranged_spell"

	None string = "none"
)

type ActionDefinition struct {
	ID          uuid.UUID
	Name        string
	Description string
	RangeNormal int
	RangeLong   int
	CostType    CostType
	AttackType  AttackType
}
