package main

type PlayerMove int

const (
	None PlayerMove = iota
	Hit
	Stay
)

type PlayerData struct {
	ID     string     `json:"id"`
	Wallet float32    `json:"wallet"`
	Bet    float32    `json:"bet"`
	Hand   Hand       `json:"hand"`
	Move   PlayerMove `json:"move"`
	Won    bool       `json:"won"`
}

var playerData PlayerData

func (p *PlayerData) Init() {
	playerData = PlayerData{
		ID:     "Jim",
		Wallet: 0,
		Bet:    0,
		Hand:   nil,
		Move:   None,
		Won:    false,
	}
}
