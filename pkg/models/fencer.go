package models

// Fencer represents Fencer information
type Fencer struct {
	ID     string `json:"fid"`
	Name   string `json:"name"`
	Weapon string `json:"weapon"`
	Rank   int    `json:"rank"`
}
