package club

import (
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/gojou/goof/pkg/models"
	"github.com/gojou/goof/pkg/services/club/repository"
	"github.com/gorilla/mux"
)

// Service todo
type Service interface {
	// GetClub(id string) (*models.Club, error)
	AddClub(w http.ResponseWriter, r *http.Request)
	ListClubs() (*[]models.Club, error)
	Serve(w http.ResponseWriter, r *http.Request)
	ServeJSON(w http.ResponseWriter, r *http.Request)
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
func (s *service) AddClub(w http.ResponseWriter, r *http.Request) {
	var club models.Club
	err := json.NewDecoder(r.Body).Decode(&club)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`"error":"Error unmarshaling the request"`))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.r.AddClub(club)
	// clist, _ := s.r.ListClubs()
	// result, _ := json.MarshalIndent(clist, "", "  ")
	result, _ := json.MarshalIndent(club, "", "  ")

	w.Write(result)
	return
}

// Serve turns Service into an HTTP server
func (s *service) Serve(w http.ResponseWriter, r *http.Request) {
	clist, _ := s.r.ListClubs()
	fmt.Fprintf(w, "%v\n", *clist)

}

// ServeJSON serves the list of fencers in JSON format
func (s *service) ServeJSON(w http.ResponseWriter, r *http.Request) {
	clist, _ := s.r.ListClubs()
	json, _ := json.MarshalIndent(clist, "", "  ")
	fmt.Fprintln(w, string(json))

}
