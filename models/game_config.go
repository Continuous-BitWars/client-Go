package models

type GameConfig struct {
	BaseLevels []BaseLevel `json:"baselevels"` // all available base levels
	Paths      PathConfig  `json:"paths"`      // settings containing paths between bases
}

