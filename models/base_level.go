package models

type BaseLevel struct {
    MaxPopulation uint32 `json:"maxPopulation"` // number of sustainable bits
    UpgradeCost uint32 `json:"upgradeCost"`     // bits required to unlock this upgrade
    SpawnRate uint32 `json:"spawnRate"`         // number uf bits spawned per tick
}