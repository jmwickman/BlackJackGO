package main

import (
	"fmt"
)

func main() {
	//initGame()
	//
	//s := CreateNewServer()
	//s.MountHandlers()
	//http.ListenAndServe(":3000", s.Router)
}

func initGame() {
	gameData.Init()
	Debug(gameData.Deck)
	gameData.Deck.Shuffle()
	Debug(gameData.Deck)
}

func Debug(d Deck) {
	for i := range d {
		fmt.Printf("Card #%d is a %s of %ss\n", i+1, d[i].Type, d[i].Suit)
	}
}
