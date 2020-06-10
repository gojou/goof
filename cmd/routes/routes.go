package routes

import (
	"github.com/gojou/goof/pkg/handlers"
	"github.com/gojou/goof/pkg/services/fencer"

	"github.com/gorilla/mux"
)

// Routing defines the API for goof
func Routing(r *mux.Router) {
	sf := fencer.NewService(r)

	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/list", sf.Serve)
	r.HandleFunc("/jlist", sf.ServeJSON)

}
