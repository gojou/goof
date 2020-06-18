package repository

import (
	"github.com/gojou/goof/pkg/models"
)

var clubs = []models.Club{
	models.Club{
		ID:   "pfc",
		Name: "Peekskill Fencing Center",
		City: "Peekskill",
	},
	models.Club{
		ID:   "chfc",
		Name: "Croton Harmon Fencing Club",
		City: "Croton-on-Hudson",
	},
}

// Repository is the foundation for the methods returned
type Repository struct {
	clubs []models.Club
}

// Club is a member of a fantasy fencing club
type Club interface {
	GetClub(id string) (*models.Club, error)
	AddClub(club models.Club) (*models.Club, error)
	ListClubs() (*[]models.Club, error)
}

//GetClub gets a club, eventually
func (*Repository) GetClub(id string) (*models.Club, error) {
	return nil, nil
}

//AddClub adds a club, eventually
func (*Repository) AddClub(club models.Club) (*models.Club, error) {
	clubs = append(clubs, club)
	return &clubs[len(clubs)-1], nil
}

//ListClubs lists all clubs, eventually
func (*Repository) ListClubs() (*[]models.Club, error) {
	return &clubs, nil
}
