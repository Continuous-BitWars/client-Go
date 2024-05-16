package models

type PlayerAction struct {
	Src    uint32 `json:"src"`    // uid of source base
	Dest   uint32 `json:"dest"`   // uid of destination base
	Amount uint32 `json:"amount"` // number of bits moved
}

