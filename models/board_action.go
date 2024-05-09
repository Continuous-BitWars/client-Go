package models

import (
	"github.com/google/uuid"
)

type BoardAction struct {
	Uuid     uuid.UUID `json:"uuid"`     // uuid of the action
	Player   uint32    `json:"player"`   // uid of the owning player
	Progress Progress  `json:"progress"` // progress of the units on the path
}
