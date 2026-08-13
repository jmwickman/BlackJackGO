package main

type PlayerData struct {
	ID     string  `json:"id"`
	Wallet float32 `json:"wallet"`
	Bet    float32 `json:"bet"`
	Hand   Hand    `json:"hand"`
	Split  bool    `json:"split"`
	Won    bool    `json:"won"`
}

var playerData PlayerData

func (p *PlayerData) Init() {
	playerData = PlayerData{
		ID:     "Jim",
		Wallet: 100.00,
		Bet:    1,
		Hand:   Hand{},
		Split:  false,
		Won:    false,
	}
}
