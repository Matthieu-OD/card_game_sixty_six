package game

type CardSuit string

const (
	Spade   CardSuit = "spade"
	Heart   CardSuit = "heart"
	Diamond CardSuit = "diamond"
	Club    CardSuit = "club"
)

type CardValue string

const (
	Eight CardValue = "8"
	Nine  CardValue = "9"
	Ten   CardValue = "10"
	Jack  CardValue = "jack"
	Queen CardValue = "queen"
	King  CardValue = "king"
	Ace   CardValue = "ace"
)

type Card struct {
	ID     int
	Suit   CardSuit
	Value  CardValue
	Points int
}
