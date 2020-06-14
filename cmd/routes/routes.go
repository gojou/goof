package routes

import (
	"github.com/gojou/goof/pkg/handlers"
	"github.com/gojou/goof/pkg/services/club"
	"github.com/gojou/goof/pkg/services/fencer"

	"github.com/gorilla/mux"
)

// Routing defines the API for goof
func Routing(r *mux.Router) {
	sf := fencer.NewService(r)
	sc := club.NewService(r)

	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/fencer/list", sf.Serve)
	r.HandleFunc("/fencer/jlist", sf.ServeJSON)
	r.HandleFunc("/club/list", sc.Serve)
	r.HandleFunc("/club/jlist", sc.ServeJSON)

}
