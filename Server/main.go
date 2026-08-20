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

func Debug(d Deck) {
	for i := range d {
		fmt.Printf("Card #%d is a %s of %ss\n", i+1, d[i].Type, d[i].Suit)
	}
}
