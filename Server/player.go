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
	p.ID = "Jim"
	p.Wallet = 100.00
	p.Bet = 0
	p.Hand = Hand{}
	p.Split = false
	p.Won = false
}
