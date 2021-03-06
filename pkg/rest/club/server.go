package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gojou/goof/pkg/models"
	"github.com/gojou/goof/pkg/services/club"
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
	s club.Service
	m *mux.Router
}

// NewServer instantiates the server
func NewServer(s club.Service, m *mux.Router) Server {
	return &server{s: s, m: m}
}

// HandleJSON serves the list of fencers in JSON format
func (s *server) HandleJSON() http.HandlerFunc {
	clubs, _ := s.s.ListClubs()
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Deal with err when not nil
		json, _ := json.MarshalIndent(clubs, "", "  ")
		fmt.Fprintln(w, string(json))
	}
}

// HandleJSON serves the list of fencers in JSON format
func (s *server) HandleHTTP() http.HandlerFunc {
	clubs, _ := s.s.ListClubs()
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Deal with err when not nil
		fmt.Fprintf(w, "%v\n", *clubs)
	}
}

func (s *server) HandleAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// method := r.Method
		// log.Printf("%v\n", method)

		if r.Method == "POST" {
			var club models.Club
			err := json.NewDecoder(r.Body).Decode(&club)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`"error":"Error unmarshaling the request"`))
				return
			}

			newclub, err := s.s.AddClub(club)
			if err != nil {
				log.Printf("%v\n", err)
				w.Write([]byte("Error adding club"))
				return
			}

			result, _ := json.MarshalIndent(newclub, "", "  ")
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
