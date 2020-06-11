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

// Fencer is a member of a fantasy fencing club
type Fencer interface {
	GetClub(id string) (*models.Club, error)
	AddClub(fencer *models.Club) (*models.Fencer, error)
	ListClubs() (*[]models.Club, error)
}

//GetClub gets a club, eventually
func (*Repository) GetClub(id string) (*models.Club, error) {
	return nil, nil
}

//AddClub adds a club, eventually
func (*Repository) AddClub(fencer *models.Club) (*models.Fencer, error) {
	return nil, nil
}

//ListClubs lists all clubs, eventually
func (*Repository) ListClubs() (*[]models.Club, error) {
	return &fencers, nil
}
