package main

/// go test ./... -v -cover

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGamePlayHandler(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	for range 1 {
		RunGame(t, s)
	}
}

func RunGame(t *testing.T, s *Server) {
	// Create new game
	req, _ := http.NewRequest("GET", "/newgame", nil)
	response := executeRequest(req, s)

	checkResponseCode(t, http.StatusOK, response.Code)

	// Dealer draw
	jsonData, _ := json.Marshal(gameData)
	req1, _ := http.NewRequest("GET", "/dealer/draw", bytes.NewBuffer(jsonData))
	response = executeRequest(req1, s)

	fmt.Println(gameData.DealerHand, gameData.DealerWon)
	fmt.Println(gameData.PlayerInfo.Hand)
	fmt.Println(gameData.PlayerInfo.Wallet, gameData.PlayerInfo.Bet)

	checkResponseCode(t, http.StatusOK, response.Code)
}
