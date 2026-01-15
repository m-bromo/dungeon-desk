package domain

type DamageType string

const (
	DamageTypeBludgeoning DamageType = "bludgeoning"
	DamageTypePiercing    DamageType = "piercing"
	DamageTypeSlashing    DamageType = "slashing"
	DamageTypeAcid        DamageType = "acid"
	DamageTypeCold        DamageType = "cold"
	DamageTypeFire        DamageType = "fire"
	DamageTypeForce       DamageType = "force"
	DamageTypeLightning   DamageType = "lightning"
	DamageTypeNecrotic    DamageType = "necrotic"
	DamageTypePoison      DamageType = "poison"
	DamageTypePsychic     DamageType = "psychic"
	DamageTypeRadiant     DamageType = "radiant"
	DamageTypeThunder     DamageType = "thunder"
)

type DamageFormula struct {
	ID         int
	DiceCount  int
	DiceValue  int
	DamageType DamageType
}
