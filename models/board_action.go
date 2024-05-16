package models

import (
	"github.com/google/uuid"
)

type BoardAction struct {
	Src      uint32    `json:"src"`
	Dest     uint32    `json:"dest"`
	Amount   uint32    `json:"amount"`
	Uuid     uuid.UUID `json:"uuid"`     // uuid of the action
	Player   uint32    `json:"player"`   // uid of the owning player
	Progress Progress  `json:"progress"` // progress of the units on the path
}
