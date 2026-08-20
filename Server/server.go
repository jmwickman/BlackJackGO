package main

import (
	"fmt"
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
	s.Router.Get("/player/bet/{bet}", PlayerBetHandler)
	s.Router.Get("/dealer/draw", DealerDrawHandler)
	s.Router.Get("/newhand",
		func(w http.ResponseWriter, r *http.Request) { NewGameHandler(w, false) })
	s.Router.Get("/newgame",
		func(w http.ResponseWriter, r *http.Request) { NewGameHandler(w, true) })
}

func NewGameHandler(w http.ResponseWriter, isNewGame bool) {

	if isNewGame {
		masterDeck = InitDeck()
		playerData.Init()
		gameData.Init(playerData)
		fmt.Println("newgame")
	}

	deck := masterDeck
	playerInfo := gameData.PlayerInfo
	newDraw := Hand{}

	// Shuffle
	deck.Shuffle()

	// Draw for the player
	newDraw, deck = deck.Draw(1)
	playerInfo.Hand = newDraw

	// Draw for the dealer
	newDraw, deck = deck.Draw(2)
	gameData.DealerHand = newDraw

	gameData.PlayerInfo = playerInfo
	gameData.Deck = deck

	WriteGameData(w, gameData)
}

func PlayerBetHandler(w http.ResponseWriter, r *http.Request) {
	bet, err := strconv.ParseFloat(chi.URLParam(r, "bet"), 32)
	if err != nil || bet < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	gameData.PlayerInfo.Bet = float32(bet)
	gameData.PlayerInfo.Wallet -= float32(bet)

	WriteGameData(w, gameData)
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

	newDraw := Hand{}

	var playResult HandEval
	var winAmount float32
	playerInfo := gameData.PlayerInfo

	if draw > 0 && draw < 3 {
		newDraw, gameData.Deck = gameData.Deck.Draw(draw)
		playerInfo.Hand = append(playerInfo.Hand, newDraw...)
		playResult, winAmount = EvaluateDrawResults(playerInfo.Hand, gameData.DealerHand, playerInfo.Bet)
	}

	gameData.PlayerInfo = playerInfo
	gameData.UpdateGameResult(playResult, winAmount)

	WriteGameData(w, gameData)
}

func DealerDrawHandler(w http.ResponseWriter, r *http.Request) {

	newDraw := Hand{}

	var playResult HandEval
	var winAmount float32
	playerInfo := gameData.PlayerInfo
	hand := gameData.DealerHand

	for playResult == NoResult {
		newDraw, gameData.Deck = gameData.Deck.Draw(1)
		hand = append(hand, newDraw...)
		playResult, winAmount = EvaluateDrawResults(playerInfo.Hand, hand, playerInfo.Bet)
	}

	gameData.PlayerInfo = playerInfo
	gameData.DealerHand = hand
	gameData.UpdateGameResult(playResult, winAmount)

	WriteGameData(w, gameData)
}
