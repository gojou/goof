package repository

import (
	"github.com/gojou/goof/pkg/models"
)

var fencers = []models.Fencer{
	models.Fencer{
		ID:     "markp",
		Name:   "Mark Poling",
		Weapon: "foil",
		Rank:   10,
	},
	models.Fencer{
		ID:     "adenp",
		Name:   "Aden Poling",
		Weapon: "foil",
		Rank:   2,
	},
	models.Fencer{
		ID:     "rhip",
		Name:   "Rhi Poling",
		Weapon: "foil",
		Rank:   2,
	},
}

// Repository is the foundation for the methods returned
type Repository struct {
	fencers []models.Fencer
}

// Fencer is a member of a fantasy fencing club
type Fencer interface {
	GetFencer(id string) (*models.Fencer, error)
	AddFencer(fencer *models.Fencer) (*models.Fencer, error)
	ListFencers() (*[]models.Fencer, error)
}

//GetFencer gets a fencer, eventually
func (*Repository) GetFencer(id string) (*models.Fencer, error) {
	return nil, nil
}

//AddFencer adds a fencer, eventually
func (*Repository) AddFencer(fencer *models.Fencer) (*models.Fencer, error) {
	return nil, nil
}

//ListFencers lists all fencers, eventually
func (*Repository) ListFencers() (*[]models.Fencer, error) {
	return &fencers, nil
}
