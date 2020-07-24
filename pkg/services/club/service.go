package club

import (
	"github.com/gojou/goof/pkg/models"
	repo "github.com/gojou/goof/pkg/services/club/firestore"
)

// Service todo
type Service interface {
	// GetClub(id string) (*models.Club, error)
	AddClub(models.Club) (*models.Club, error)
	ListClubs() (*[]models.Club, error)
}

type service struct {
	r *repo.Repository
}

// NewService initializes the dependencies for the Club service
func NewService() Service {
	r := new(repo.Repository)
	return &service{r}
}

// ListClubs returns a list of all fencers in the repo
func (s *service) ListClubs() (*[]models.Club, error) {
	return s.r.ListClubs()
}

// AddClub adds a club to the repository
func (s *service) AddClub(club models.Club) (newclub *models.Club, err error) {
	s.r.AddClub(club)
	return newclub, err
}
