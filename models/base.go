package models

type Base struct {
	Name              string   `json:"name"`
	Position          Position `json:"position"`
	Uid               uint32   `json:"uid"`
	Player            uint32   `json:"player"`
	Population        uint32   `json:"population"`
	Level             uint32   `json:"level"`
	UnitsUntilUpgrade uint32   `json:"units_until_upgrade"`
}
