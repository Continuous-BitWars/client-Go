package models

type GameConfig struct {
    BaseLevels []BaseLevel `json:"baseLevel"` // all available base levels
    Paths PathConfig `json:"pathConfig"`      // settings containing paths between bases
}