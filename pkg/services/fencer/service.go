package fencer

import (
	"github.com/gojou/goof/pkg/models"
	"github.com/gojou/goof/pkg/services/fencer/repository"
)

// service adds the local dependencies
type service struct {
	r *repository.Repository
}

// Service todo
type Service interface {
	GetFencer(id string) (*models.Fencer, error)
	AddFencer(fencer *models.Fencer) (*models.Fencer, error)
	ListFencers() (*[]models.Fencer, error)
}

// NewService initializes the fencer service
func NewService() Service {
	r := new(repository.Repository)

	return &service{r}
}

// GetFencer returns a fencer struct from the repository
func (*service) GetFencer(id string) (*models.Fencer, error) {
	return nil, nil
}

// AddFencer adds a fencer struct to the repository
func (*service) AddFencer(fencer *models.Fencer) (*models.Fencer, error) {
	return nil, nil
}

// ListFencers returns a list of all fencers in the repo
func (s *service) ListFencers() (*[]models.Fencer, error) {
	return s.r.ListFencers()
}
