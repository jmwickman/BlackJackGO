package main

type GameData struct {
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
		Deck:       InitDeck(),
		DealerHand: Hand{},
		PlayerInfo: playerData,
		GameOver:   false,
		DealerWon:  false,
	}
}
