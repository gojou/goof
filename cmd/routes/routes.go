package routes

import (
	"github.com/gojou/goof/pkg/handlers"
	//Alias the club and fencer REST services
	rclub "github.com/gojou/goof/pkg/rest/club"
	rfencer "github.com/gojou/goof/pkg/rest/fencer"
	"github.com/gojou/goof/pkg/services/club"
	"github.com/gojou/goof/pkg/services/fencer"

	"github.com/gorilla/mux"
)

// Routing defines the API for goof
func Routing(r *mux.Router) {
	sf := fencer.NewService()
	sc := club.NewService(r)
	sfh := rfencer.NewServer(sf, r)
	sch := rclub.NewServer(sc, r)

	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/club/list", sch.HandleHTTP())
	r.HandleFunc("/club/jlist", sch.HandleJSON())
	r.HandleFunc("/fencer/list", sfh.HandleHTTP())
	r.HandleFunc("/fencer/jlist", sfh.HandleJSON())
	r.HandleFunc("/club/add", sch.HandleAdd())

	// r.HandleFunc("/club", sc.AddClub).Methods("POST")
	// r.HandleFunc("/club/list", sc.Serve).Methods("GET")
	// r.HandleFunc("/club/jlist", sc.ServeJSON).Methods("GET")

}
