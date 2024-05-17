package models

type PathConfig struct {
	GracePeriod uint32 `json:"grace_period"` // time until groups of bits take damage
	DeathRate   uint32 `json:"death_rate"`   // number of units killed every tick after surpassing the grace period
}

