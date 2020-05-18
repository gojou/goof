package repository

import "github.com/gojou/goof/pkg/models"

//
// Repository is the foundation for the methods returned
type Repository struct {
	fencers []models.Fencer
}

// Fencer is a member of a fantasy fencing club
type Fencer interface {
	GetFencer(id string) (*models.Fencer, error)
	AddFencer(fencer *models.Fencer) (*models.Fencer, error)
	ListFencers() ([]models.Fencer, error)
}

//GetFencer gets a fencer, eventually
func (*Repository) GetFencer(id string) (*models.Fencer, error) {
	return nil, nil
}

//AddFencer adds a fencer, eventually
func (*Repository) AddFencer(fencer *models.Fencer) (*models.Fencer, error) {
	return nil, nil
}

//ListFencers adds a fencer, eventually
func (*Repository) ListFencers() ([]models.Fencer, error) {
	return nil, nil
}
