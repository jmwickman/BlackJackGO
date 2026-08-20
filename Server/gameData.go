package main

type GameData struct {
	Deck       Deck       `json:"deck"`
	DealerHand Hand       `json:"dealer_hand"`
	PlayerInfo PlayerData `json:"player_info"`
	GameOver   bool       `json:"game_over"`
	DealerWon  bool       `json:"dealer_won"`
}

var gameData GameData

func (g *GameData) Init(p PlayerData) {
	g.Deck = Deck{}
	g.DealerHand = Hand{}
	g.PlayerInfo = p
	g.GameOver = false
	g.DealerWon = false
}

func (g *GameData) UpdateGameResult(h HandEval, w float32) {
	if w > 0 {
		g.PlayerInfo.Wallet += w
	}
	g.PlayerInfo.Won = h == PlayerWin
	g.DealerWon = h == DealerWin
	g.GameOver = h != NoResult
}
