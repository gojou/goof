package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gojou/goof/pkg/models"
	"github.com/gojou/goof/pkg/services/fencer"
	"github.com/gorilla/mux"
)

// Server defines the methods on the server
type Server interface {
	ServeJSON(*[]models.Fencer, error) http.HandlerFunc
}

type server struct {
	s *fencer.Service
	m *mux.Router
}

// NewServer instantiates the server
func NewServer(s *fencer.Service, m *mux.Router) Server {
	return &server{s: s, m: m}
}

// ServeJSON serves the list of fencers in JSON format
func (s *server) ServeJSON(fencers *[]models.Fencer, err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		json, _ := json.MarshalIndent(fencers, "", "  ")
		fmt.Fprintln(w, string(json))
	}
}
