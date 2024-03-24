package models

type BoardAction struct {
    Uuid uuid `json:"uuid"`             // uuid of the action
    Player uint32 `json:"player"`       // uid of the owning player
    Progress Progress `json:"progress"` // progress of the units on the path
}