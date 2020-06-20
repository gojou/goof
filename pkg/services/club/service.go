package club

import (
	"github.com/gojou/goof/pkg/models"
	"github.com/gojou/goof/pkg/services/club/repository"
	"github.com/gorilla/mux"
)

// Service todo
type Service interface {
	// GetClub(id string) (*models.Club, error)
	AddClub(models.Club) (*models.Club, error)
	ListClubs() (*[]models.Club, error)
	// Serve(w http.ResponseWriter, r *http.Request)
	// ServeJSON(w http.ResponseWriter, r *http.Request)
}

type service struct {
	r *repository.Repository
	m *mux.Router
}

// NewService todo
func NewService(mux *mux.Router) Service {
	r := new(repository.Repository)

	return &service{r, mux}
}

// ListFencers returns a list of all fencers in the repo
func (s *service) ListClubs() (*[]models.Club, error) {
	return s.r.ListClubs()
}

// AddClub will add a club to the repository
func (s *service) AddClub(club models.Club) (newclub *models.Club, err error) {
	s.r.AddClub(club)
	return newclub, err
}

// // Serve turns Service into an HTTP server
// func (s *service) Serve(w http.ResponseWriter, r *http.Request) {
// 	clist, _ := s.r.ListClubs()
// 	fmt.Fprintf(w, "%v\n", *clist)
//
// }
//
// // ServeJSON serves the list of fencers in JSON format
// func (s *service) ServeJSON(w http.ResponseWriter, r *http.Request) {
// 	clist, _ := s.r.ListClubs()
// 	json, _ := json.MarshalIndent(clist, "", "  ")
// 	fmt.Fprintln(w, string(json))
//
// }
