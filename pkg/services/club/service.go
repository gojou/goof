package club

import (
	repo "github.com/gojou/goof/pkg/services/club/firestore"

	"github.com/gorilla/mux"
)

// Service todo
type Service interface {
	// GetClub(id string) (*models.Club, error)
	AddClub(Club) (*Club, error)
	ListClubs() (*[]Club, error)
}

type service struct {
	r *repo.Repository
	m *mux.Router
}

// NewService todo
func NewService(mux *mux.Router) Service {
	r := new(repo.Repository)

	return &service{r, mux}
}

// ListFencers returns a list of all fencers in the repo
func (s *service) ListClubs() (*[]Club, error) {
	return s.r.ListClubs()
}

// AddClub will add a club to the repository
func (s *service) AddClub(club Club) (newclub *Club, err error) {
	s.r.AddClub(club)
	return newclub, err
}
