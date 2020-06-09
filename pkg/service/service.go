package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gojou/goof/pkg/models"
	"github.com/gojou/goof/pkg/repository"
	"github.com/gorilla/mux"
)

// service is just this struct, you know?
type service struct {
	r *repository.Repository
	m *mux.Router
}

// Service todo
type Service interface {
	GetFencer(id string) (*models.Fencer, error)
	AddFencer(fencer *models.Fencer) (*models.Fencer, error)
	ListFencers() (*[]models.Fencer, error)
	Serve(w http.ResponseWriter, r *http.Request)
	ServeJSON(w http.ResponseWriter, r *http.Request)
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

// ListFencers returns a list of all fencers in the repo
func (s *service) ListFencers() (*[]models.Fencer, error) {
	return s.r.ListFencers()
}

// Serve turns Service into an HTTP server
func (s *service) Serve(w http.ResponseWriter, r *http.Request) {
	flist, _ := s.r.ListFencers()
	fmt.Fprintf(w, "%v\n", *flist)

}

// ServeJSON serves the list of fencers in JSON format
func (s *service) ServeJSON(w http.ResponseWriter, r *http.Request) {
	flist, _ := s.r.ListFencers()
	json, _ := json.MarshalIndent(flist, "", "  ")
	fmt.Fprintln(w, string(json))

}
