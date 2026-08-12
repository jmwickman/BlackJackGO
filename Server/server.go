package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {
	// Mount all Middleware here
	s.Router.Use(middleware.Logger)

	// Mount all handlers here
	s.Router.Get("/player/draw/{count}",
		func(w http.ResponseWriter, r *http.Request) { DrawHandler(w, r, true) })
	s.Router.Get("/dealer/draw/{count}", func(w http.ResponseWriter, r *http.Request) { DrawHandler(w, r, false) })

}
func DrawHandler(w http.ResponseWriter, r *http.Request, isPlayer bool) {

	draw, err := strconv.Atoi(chi.URLParam(r, "count"))
	if err != nil {
		log.Fatal(err)
	}

	if draw > 0 {
		data := GameData{}
		newDraw := Hand{}

		err1 := json.NewDecoder(r.Body).Decode(&data)
		if err1 != nil {
			log.Fatal("Error reading Draw Card request", err1)
		}

		newDraw, data.Deck = data.Deck.Draw(draw)

		if isPlayer {
			data.PlayerInfo.Hand = append(data.PlayerInfo.Hand, newDraw...)
		} else {
			data.DealerHand = append(data.DealerHand, newDraw...)
		}

		playResult := EvaluateDrawResults(data.PlayerInfo.Hand, data.DealerHand)

		data.PlayerInfo.Won = playResult == PlayerWin
		data.DealerWon = playResult == DealerWin
		data.GameOver = playResult != NoResult

		gameData = data

		// Make the response
		w.Header().Set("Content-Type", "application/json")

		err2 := json.NewEncoder(w).Encode(gameData)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
