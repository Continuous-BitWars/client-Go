package models

type GameState struct {
	Actions []BoardAction `json:"actions"`
	Bases   []Base        `json:"bases"`
	Config  GameConfig    `json:"config"`
	Game    Game          `json:"game"`
}
