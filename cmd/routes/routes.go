package routes

import (
	"github.com/gojou/goof/pkg/handlers"
	"github.com/gojou/goof/pkg/rest"
	"github.com/gojou/goof/pkg/services/club"
	"github.com/gojou/goof/pkg/services/fencer"

	"github.com/gorilla/mux"
)

// Routing defines the API for goof
func Routing(r *mux.Router) {
	sf := fencer.NewService(r)
	sc := club.NewService(r)
	sh := rest.NewServer(&sf, r)

	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/fencer/list", sh.HandleHTTP(sf.ListFencers())).Methods("GET")
	r.HandleFunc("/fencer/jlist", sh.HandleJSON(sf.ListFencers())).Methods("GET")
	r.HandleFunc("/club", sc.AddClub).Methods("POST")
	r.HandleFunc("/club/list", sc.Serve).Methods("GET")
	r.HandleFunc("/club/jlist", sc.ServeJSON).Methods("GET")

}
