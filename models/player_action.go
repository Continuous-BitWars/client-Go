package models

type PlayerAction struct {
    src uint32 `json:"distance"`    // uid of source base
    dest uint32 `json:"distance"`   // uid of destination base
    amount uint32 `json:"distance"` // number of bits moved
}