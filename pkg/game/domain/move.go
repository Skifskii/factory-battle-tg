package domain

type CardName string

const (
	CardIncrease CardName = "inc"
	CardDecrease CardName = "dec"
)

// Move - ход.
type Move struct {
	PlayerID int64
	CardName CardName
}
