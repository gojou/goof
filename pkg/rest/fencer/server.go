package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gojou/goof/pkg/models"
	"github.com/gojou/goof/pkg/services/fencer"
	"github.com/gorilla/mux"
)

// Server defines the methods on the server
type Server interface {
	HandleJSON() http.HandlerFunc
	HandleHTTP() http.HandlerFunc
	HandleAdd() http.HandlerFunc
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

func (s *server) HandleAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// method := r.Method
		// log.Printf("%v\n", method)

		if r.Method == "POST" {
			var fencer models.Fencer
			err := json.NewDecoder(r.Body).Decode(&fencer)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`"error":"Error unmarshaling the request"`))
				return
			}

			newfencer, err := s.s.AddFencer(fencer)
			if err != nil {
				log.Printf("%v\n", err)
				w.Write([]byte("Error adding fencer"))
				return
			}

			result, _ := json.MarshalIndent(newfencer, "", "  ")
			log.Printf("Output:\n%v\n", result)

			w.WriteHeader(http.StatusOK)
			w.Write(result)
			return

		}
		log.Printf("%v\n", r.Method)
		w.Write([]byte("It's a \"GET\""))
		return
	}
}
