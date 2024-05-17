package models

type BaseLevel struct {
	MaxPopulation uint32 `json:"max_population"` // number of sustainable bits
	UpgradeCost   uint32 `json:"upgrade_cost"`   // bits required to unlock this upgrade
	SpawnRate     uint32 `json:"spawn_rate"`     // number uf bits spawned per tick
}

