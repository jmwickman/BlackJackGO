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
	s.Router.Get("/dealer/draw/{firstDeal}",
		func(w http.ResponseWriter, r *http.Request) { DrawHandler(w, r, false) })

}
func DrawHandler(w http.ResponseWriter, r *http.Request, isPlayer bool) {

	var playerDraw = 0
	var firstDeal = false
	if isPlayer {
		draw, err := strconv.Atoi(chi.URLParam(r, "count"))
		if err != nil || draw < 1 {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}

		if draw == 0 {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Draw value less than 1.")
		}

		playerDraw = draw
	} else {
		deal, err := strconv.ParseBool(chi.URLParam(r, "firstDeal"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}

		firstDeal = deal
	}

	data := GameData{}
	newDraw := Hand{}

	err1 := json.NewDecoder(r.Body).Decode(&data)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err1)
	}

	var playResult HandEval

	if isPlayer {

		if playerDraw > 0 && playerDraw < 3 {
			newDraw, data.Deck = data.Deck.Draw(playerDraw)
			data.PlayerInfo.Hand = append(data.PlayerInfo.Hand, newDraw...)
			playResult = EvaluateDrawResults(data.PlayerInfo.Hand, data.DealerHand)
		}

	} else {

		for playResult == NoResult {

			drawCount := 1
			if firstDeal {
				drawCount = 2
			}

			newDraw, data.Deck = data.Deck.Draw(drawCount)
			data.DealerHand = append(data.DealerHand, newDraw...)
			playResult = EvaluateDrawResults(data.PlayerInfo.Hand, data.DealerHand)
		}
	}

	data.PlayerInfo.Won = playResult == PlayerWin
	data.DealerWon = playResult == DealerWin
	data.GameOver = playResult != NoResult

	gameData = data

	// Make the response
	w.Header().Set("Content-Type", "application/json")

	err2 := json.NewEncoder(w).Encode(gameData)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err2)
	}
}
