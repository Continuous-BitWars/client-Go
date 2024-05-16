package logic

import "github.com/Continuous-BitWars/player-Go/models"

func Decide(gameState models.GameState) []models.PlayerAction {
	// TODO: place your player logic here.

	var actions = []models.PlayerAction{
		{
			Src:    0,
			Dest:   0,
			Amount: 0,
		},
	}
	return actions
}
