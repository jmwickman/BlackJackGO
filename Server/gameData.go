package main

type GameData struct {
	Pot        float32    `json:"pot"`
	Deck       Deck       `json:"deck"`
	DealerHand Hand       `json:"dealer_hand"`
	PlayerInfo PlayerData `json:"player_info"`
	GameOver   bool       `json:"game_over"`
	DealerWon  bool       `json:"dealer_won"`
}

var gameData GameData

func (g *GameData) Init() {
	playerData.Init()
	gameData = GameData{
		Pot:        0,
		Deck:       InitDeck(),
		DealerHand: nil,
		PlayerInfo: playerData,
		GameOver:   false,
		DealerWon:  false,
	}
}
