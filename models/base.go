package models

type Base struct {
    Position Position `json:"position"`   // position of the base
    Uid uint32 `json:"uid"`               // uid of the base
    Player uint32 `json:"player"`         // uid of the owning player
    Population uint32 `json:"population"` // number of bits in the base
    Level uint32 `json:"level"`           // level of the base
    UnitsUntilUpgrade uint32 `json:"bet"` // bits needed for until the next upgrade
}