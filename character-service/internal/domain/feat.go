package domain

type Feat struct {
	ID               int
	Name             string
	Description      string
	ActionDefinition *ActionDefinition
}
