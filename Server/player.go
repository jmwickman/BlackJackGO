package main

type playerData struct {
	Wallet float32 `json:"wallet"`
	Hand   []Card  `json:"hand"`
}
