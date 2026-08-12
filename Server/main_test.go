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

func TestDrawHandler(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	testData := GameData{
		Pot:        0,
		Deck:       InitDeck(),
		DealerHand: Hand{},
		PlayerInfo: PlayerData{
			ID:     "Jim",
			Wallet: 0,
			Bet:    0,
			Hand:   Hand{},
			Move:   None,
			Won:    false,
		},
		GameOver:  false,
		DealerWon: false,
	}

	jsonData, _ := json.Marshal(testData)

	req, _ := http.NewRequest("GET", "/player/draw/1", bytes.NewBuffer(jsonData))

	response := executeRequest(req, s)

	checkResponseCode(t, http.StatusOK, response.Code)

	fmt.Println(response.Body.String())
}
