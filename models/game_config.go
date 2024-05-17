package models

type GameConfig struct {
	BaseLevels []BaseLevel `json:"base_levels"` // all available base levels
	Paths      PathConfig  `json:"paths"`       // settings containing paths between bases
}
