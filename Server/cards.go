package main

import (
	"math/rand"
	"time"
)

type Card struct {
	Type string `json:"type"`
	Suit string `json:"suit"`
}

type Deck []Card

const (
	Ace   = "ace"
	Two   = "two"
	Three = "three"
	Four  = "four"
	Five  = "five"
	Six   = "six"
	Seven = "seven"
	Eight = "eight"
	Nine  = "nine"
	Ten   = "ten"
	Jack  = "jack"
	Queen = "queen"
	King  = "king"
)

const (
	Heart   = "heart"
	Diamond = "diamond"
	Club    = "club"
	Spade   = "spade"
)

var cardValues = map[string]int{
	Ace:   11,
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   10,
	Jack:  10,
	Queen: 10,
	King:  10,
}

func InitDeck() (deck Deck) {
	types := []string{Ace, Two, Three, Four, Five, Six,
		Seven, Eight, Nine, Ten, Jack, Queen, King}

	suits := []string{Heart, Diamond, Club, Spade}
	
	for i := range types {
		for n := range suits {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}
	return
}

func Shuffle(d Deck) Deck {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	return d
}

func ParseCard(card string) int {
	return cardValues[card]
}
