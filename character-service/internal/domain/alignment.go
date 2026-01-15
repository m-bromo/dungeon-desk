package domain

type Alignment string

const (
	LawfulGood     Alignment = "lawful_good"
	NeutralGood    Alignment = "neutral_good"
	ChaoticGood    Alignment = "chaotic_good"
	LawfulNeutral  Alignment = "lawful_neutral"
	TrueNeutral    Alignment = "true_neutral"
	ChaoticNeutral Alignment = "chaotic_neutral"
	LawfulEvil     Alignment = "lawful_evil"
	NeutralEvil    Alignment = "neutral_evil"
	ChaoticEvil    Alignment = "chaotic_evil"
)
