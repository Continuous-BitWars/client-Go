package models

type GameState struct {
	Actions []BoardAction `json:"actions"`
	bases   []Base        `json:"bases"`
	config  []GameConfig  `json:"config"`
	game    Game          `json:"game"`
}
