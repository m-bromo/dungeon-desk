package domain

type Spell struct {
	ID               int
	Name             string
	School           SpellSchool
	ActionDefinition ActionDefinition
}
