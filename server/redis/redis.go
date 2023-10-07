package redidsclient

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return c
}

type GameID string

type Player struct {
}

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

/*
Struct of the game in redis

The goal of redis is to have one true state for all the games

need a card representation

gameId id string -- could be the key to the map {
player1 player1Cards map
player2 player1Cards map
stack
lastFold
asset
totalScore
gameScore
last round winner
}

playerCards {
hand
won
playing

}

*/

// TODO: add function to add a new party data
// TODO: add function to get the party data
