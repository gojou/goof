package service

import (
	"github.com/gojou/goof/pkg/models"
	"github.com/gojou/goof/pkg/repository"
	"github.com/gorilla/mux"
)

// Service is just this struct, you know?

type service struct {
	r *repository.Repository
	m *mux.Router
}

// Service todo
type Service interface {
	GetFencer(id string) (*models.Fencer, error)
	AddFencer(fencer *models.Fencer) (*models.Fencer, error)
	ListFencers() ([]models.Fencer, error)
}

// NewService todo
func NewService(mux *mux.Router) Service {
	r := new(repository.Repository)

	return &service{r, mux}
}

// GetFencer todo
func (*service) GetFencer(id string) (*models.Fencer, error) {
	return nil, nil
}

// AddFencer todo
func (*service) AddFencer(fencer *models.Fencer) (*models.Fencer, error) {
	return nil, nil
}

// ListFencers todo
func (*service) ListFencers() ([]models.Fencer, error) {
	return nil, nil
}
