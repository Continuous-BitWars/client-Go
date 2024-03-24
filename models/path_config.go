package models

type PathConfig struct {
    GracePeriod uint32 `json:"gracePeriod"` // time until groups of bits take damage
    DeathRate uint32 `json:"deathRate"`     // number of units killed every tick after surpassing the grace period
}