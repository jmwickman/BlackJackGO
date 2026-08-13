package main

import (
	"math/rand"
	"time"
)

type Deck []Card
type Hand []Card

func InitDeck() (d Deck) {
	types := []string{Ace, Two, Three, Four, Five, Six,
		Seven, Eight, Nine, Ten, Jack, Queen, King}

	suits := []string{Heart, Diamond, Club, Spade}

	for i := range types {
		for n := range suits {
			card := Card{
				Type:  types[i],
				Suit:  suits[n],
				Value: cardValues[types[i]],
			}
			d = append(d, card)
		}
	}
	return
}

func (d Deck) Shuffle() {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func (d Deck) Draw(count int) (Hand, Deck) {
	i := len(d) - count
	return Hand(d[i:]), d[:i]
}
