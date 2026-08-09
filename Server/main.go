package main

import (
	"fmt"
)

func main() {
	//r := chi.NewRouter()
	//r.Use(middleware.Logger)

	play()

	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello World!"))
	//})
	//http.ListenAndServe(":3000", r)
}

func play() {
	deck := InitDeck()
	Debug(deck)
	Shuffle(deck)
	Debug(deck)
}

func Debug(d Deck) {
	for i := 0; i < len(d); i++ {
		fmt.Printf("Card #%d is a %s of %ss\n", i+1, d[i].Type, d[i].Suit)
	}
}
