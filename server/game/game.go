package game

import (
	"encoding/json"
)

type GameID string

type Game struct {
	GameID       GameID
	Player2Ready bool

	Player1Hand [6]Card
	Player2Hand [6]Card

	PlayingCards [2]Card
	Stack        []Card
	Asset        Card
	LastFold     [2]Card

	Player1Score int
	Player2Score int
	TotalScore   [2]int
}

func gameToHash(g Game) (map[string]interface{}, error) {
	gameHash := make(map[string]interface{})

	gameHash["GameID"] = g.GameID
	gameHash["Player2Ready"] = g.Player2Ready

	gameHash["Player1Hand"] = stackToString(g.Player1Hand)

	gameHash["TotalScore"] = g.TotalScore
	gameHash["GameScore"] = g.GameScore
	gameHash["CurrentStage"] = g.CurrentStage
	gameHash["LastRoundWinner"] = g.LastRoundWinner

	return gameHash, nil
}
