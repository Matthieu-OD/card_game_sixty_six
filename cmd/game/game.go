package game

type Game struct {
	GameID        string
	OpponentReady bool

	Player1Cards []Card
	Player2Cards []Card

	PlayingCards []Card
	Stack        []Card
	Asset        *Card
	LastFold     []Card

	RoundScore1 int
	TotalScore1 int
	RoundScore2 int
	TotalScore2 int

	Turn int
}
