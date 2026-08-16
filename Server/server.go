package main

import (
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
	s.Router.Get("/player/draw/{count}", PlayerDrawHandler)
	s.Router.Get("/dealer/draw", DealerDrawHandler)
	s.Router.Get("/newhand",
		func(w http.ResponseWriter, r *http.Request) { NewGameHandler(w, r, false) })
	s.Router.Get("/newgame",
		func(w http.ResponseWriter, r *http.Request) { NewGameHandler(w, r, true) })
}

func PlayerDrawHandler(w http.ResponseWriter, r *http.Request) {

	draw, err := strconv.Atoi(chi.URLParam(r, "count"))
	if err != nil || draw < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	if draw == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Draw value less than 1.")
	}

	data := ReadGameData(w, r)
	newDraw := Hand{}

	var playResult HandEval
	var winAmount float32
	playerInfo := data.PlayerInfo

	if draw > 0 && draw < 3 {
		newDraw, data.Deck = data.Deck.Draw(draw)
		playerInfo.Hand = append(playerInfo.Hand, newDraw...)
		playResult, winAmount = EvaluateDrawResults(playerInfo.Hand, data.DealerHand, playerInfo.Bet)
	}

	if winAmount > 0 {
		playerInfo.Wallet += winAmount
	}
	data.PlayerInfo = playerInfo
	data.PlayerInfo.Won = playResult == PlayerWin
	data.DealerWon = playResult == DealerWin
	data.GameOver = playResult != NoResult

	gameData = data

	WriteGameData(w, gameData)
}

func DealerDrawHandler(w http.ResponseWriter, r *http.Request) {

	data := ReadGameData(w, r)
	newDraw := Hand{}

	var playResult HandEval
	var winAmount float32
	playerInfo := data.PlayerInfo

	for playResult == NoResult {
		newDraw, data.Deck = data.Deck.Draw(1)
		data.DealerHand = append(data.DealerHand, newDraw...)
		playResult, winAmount = EvaluateDrawResults(playerInfo.Hand, data.DealerHand, playerInfo.Bet)
	}

	if winAmount > 0 {
		playerInfo.Wallet += winAmount
	}

	data.PlayerInfo = playerInfo
	data.PlayerInfo.Won = playResult == PlayerWin
	data.DealerWon = playResult == DealerWin
	data.GameOver = true

	gameData = data

	WriteGameData(w, gameData)
}

func NewGameHandler(w http.ResponseWriter, r *http.Request, isNewGame bool) {

	var data GameData
	if isNewGame {
		gameData.Init()
		data = gameData
	} else {
		data = ReadGameData(w, r)
		data.Deck = InitDeck()
	}

	data.Deck.Shuffle()
	newDraw := Hand{}

	// Draw for the player
	newDraw, data.Deck = data.Deck.Draw(2)
	data.PlayerInfo.Hand = newDraw

	// Draw for the dealer
	newDraw, data.Deck = data.Deck.Draw(2)
	data.DealerHand = newDraw

	gameData = data

	WriteGameData(w, gameData)
}
