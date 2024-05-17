package models

type Game struct {
	Uid              uint32 `json:"uid"`               // uid of game
	Tick             uint32 `json:"tick"`              // tick in game
	PlayerCount      uint32 `json:"player_count"`      // number of players
	RemainingPlayers uint32 `json:"remaining_players"` // number of players remaining
	Player           uint32 `json:"player"`            // uid of your player
}

