package models

// Fencer represents Fencer information
type Fencer struct {
	ID     string `json:"id" firestore:"id"`
	Name   string `json:"name" firestore:"name"`
	Weapon string `json:"weapon" firestore:"weapon"`
	Rank   int    `json:"rank" firestore:"rank"`
}
