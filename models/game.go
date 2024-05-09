package models

type Game struct {
    Uid uint32 `json:"uid"`                           // uid of game
    Tick uint32 `json:"tick"`                         // tick in game
    PlayerCount uint32 `json:"playerCount"`           // number of players
    RemainingPlayers uint32 `json:"remainingPlayers"` // number of players remaining
    Player uint32 `json:"player"`                     // uid of your player
}