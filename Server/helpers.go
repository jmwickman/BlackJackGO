package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

func IsNilOrZero(a any) bool {
	return reflect.ValueOf(a).IsNil() || reflect.ValueOf(a).IsZero()
}

func GetHandValue(h Hand) int {

	if IsNilOrZero(h) {
		return 0
	}

	totalValue := 0

	for _, c := range h {
		totalValue += c.Value
	}

	// If we're gonna bust with an Ace, change the value.
	if totalValue > 21 {
		for i := range h {
			if h[i].Value == 11 {
				h[i].Value = 1
				totalValue -= 10
				break
			}
		}
	}

	return totalValue
}

func ReadGameData(w http.ResponseWriter, r *http.Request) GameData {

	data := GameData{}

	err1 := json.NewDecoder(r.Body).Decode(&data)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err1)
	}

	return data
}

func WriteGameData(w http.ResponseWriter, d GameData) {

	// Make the response
	w.Header().Set("Content-Type", "application/json")

	err2 := json.NewEncoder(w).Encode(d)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err2)
	}
}
