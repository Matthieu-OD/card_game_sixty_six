package game

type GameID string

type Game struct {
	GameID        GameID
	OpponentReady bool

	Player1Hand []Card
	Player2Hand []Card

	PlayingCards []Card
	Stack        []Card
	Asset        Card
	LastFold     []Card

	GameScore1  int
	TotalScore1 int
	GameScore2  int
	TotalScore2 int

	Turn int
}

func (g Game) GameToHash() (map[string]interface{}, error) {
	gameHash := make(map[string]interface{})

	gameHash["GameID"] = g.GameID
	gameHash["OpponentReady"] = g.OpponentReady

	gameHash["Player1Hand"] = stackToStringStack(g.Player1Hand)
	gameHash["Player2Hand"] = stackToStringStack(g.Player2Hand)

	gameHash["PlayingCards"] = stackToStringStack(g.PlayingCards)
	gameHash["Stack"] = stackToStringStack(g.Stack)
	gameHash["Asset"] = g.Asset.toString()
	gameHash["LastFold"] = stackToStringStack(g.LastFold)

	gameHash["TotalScore1"] = g.TotalScore1
	gameHash["TotalScore2"] = g.TotalScore2
	gameHash["GameScore1"] = g.GameScore1
	gameHash["GameScore2"] = g.GameScore2
	gameHash["Turn"] = g.Turn

	return gameHash, nil
}
