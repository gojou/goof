package models

// Fencer is the represents Fencer information
type Fencer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Weapon string `json:"weapon"`
	Rank   int    `json:"rank"`
}
