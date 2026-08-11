package main

// curl -v localhost:3000 -d 'hello world'

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"strconv"
)

func registerRoutes() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/draw/{count}", drawHandler)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}

	return router
}

func drawHandler(w http.ResponseWriter, r *http.Request) {

	data := GameData{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Fatal("Error reading Draw Card request", err)
	}

	draw, err1 := strconv.Atoi(chi.URLParam(r, "count"))
	if err1 != nil {
		log.Fatal(err)
	}

	data.PlayerInfo.Hand, data.Deck = data.Deck.Draw(draw)
	playResult := EvaluateDrawResults(data.PlayerInfo.Hand, data.DealerHand)

	data.PlayerInfo.Won = playResult == PlayerWin
	data.DealerWon = playResult == DealerWin
	data.GameOver = playResult != NoResult

	gameData = data

	// Make the response
	w.Header().Set("Content-Type", "application/json")

	err2 := json.NewEncoder(w).Encode(gameData)
	if err2 != nil {
		log.Fatal(err)
	}
}
