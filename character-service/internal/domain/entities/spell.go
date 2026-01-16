package entities

type Spell struct {
	ID               int
	Name             string
	School           SpellSchool
	ActionDefinition ActionDefinition
}
