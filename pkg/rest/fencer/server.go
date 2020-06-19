package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gojou/goof/pkg/services/fencer"
	"github.com/gorilla/mux"
)

// Server defines the methods on the server
type Server interface {
	HandleJSON() http.HandlerFunc
	HandleHTTP() http.HandlerFunc
}

// server defines the local dependencies
type server struct {
	s fencer.Service
	m *mux.Router
}

// NewServer instantiates the server
func NewServer(s fencer.Service, m *mux.Router) Server {
	return &server{s: s, m: m}
}

// HandleJSON serves the list of fencers in JSON format
func (s *server) HandleJSON() http.HandlerFunc {
	fencers, _ := s.s.ListFencers()

	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Deal with err when not nil
		json, _ := json.MarshalIndent(fencers, "", "  ")
		fmt.Fprintln(w, string(json))
	}
}

// HandleJSON serves the list of fencers in JSON format
func (s *server) HandleHTTP() http.HandlerFunc {
	fencers, _ := s.s.ListFencers()

	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Deal with err when not nil
		fmt.Fprintf(w, "%v\n", *fencers)
	}
}
