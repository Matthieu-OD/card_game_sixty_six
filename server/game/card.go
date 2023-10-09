package game

import (
	"log"
	"strconv"
)

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

var fromStringToCardSuit = map[string]CardSuit{
	"s": Spade,
	"h": Heart,
	"d": Diamond,
	"c": Club,
}

var fromStringToCardValue = map[string]CardValue{
	"8": Eight,
	"9": Nine,
	"1": Ten,
	"j": Jack,
	"q": Queen,
	"k": King,
	"a": Ace,
}

// Simplify the data stored in the redis db
func (c Card) toString() string {
	cstr := string(c.Suit[0]) + string(c.Value[0]) + strconv.Itoa(c.Points)
	return cstr
}

func toCard(stringCard string) Card {
	score, err := strconv.Atoi(stringCard[2:])
	if err != nil {
		log.Fatal(err)
	}
	return Card{
		Suit:   fromStringToCardSuit[string(stringCard[0])],
		Value:  fromStringToCardValue[string(stringCard[1])],
		Points: score,
	}
}

func stackToStringStack(stack []Card) []string {
	var stackString []string
	for _, c := range stack {
		stackString = append(stackString, c.toString())
	}
	return stackString
}

func stringStackToStack(stringStack []string) []Card {
	var stack []Card
	for _, c := range stringStack {
		stack = append(stack, toCard(c))
	}
	return stack
}
