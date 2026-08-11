package main

type Card struct {
	Type  string `json:"type"`
	Suit  string `json:"suit"`
	Value int    `json:"value"`
}

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
