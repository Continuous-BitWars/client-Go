package models

type Progress struct {
    Distance uint32 `json:"distance"` // distance between the bases
    Traveled uint32 `json:"traveled"` // distance already traveled
}