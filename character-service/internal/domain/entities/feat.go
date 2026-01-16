package entities

type Feat struct {
	ID               int
	Name             string
	Description      string
	ActionDefinition *ActionDefinition
}
