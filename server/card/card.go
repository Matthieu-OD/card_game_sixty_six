package card

type GameID string

type Player struct {
	hand    [6]Card
	won     []Card
	playing *Card
}

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
	Suit   CardSuit
	Value  CardValue
	Points int
}

type Game struct {
	Player1         Player
	Player2         Player
	Stack           []Card
	Asset           Card
	LastFold        [2]Card
	TotalScore      [2]int
	GameScore       [2]int
	LastRoundWinner string
	CurrentStage    int
}
