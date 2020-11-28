package routes

import (
	"github.com/gojou/goof/pkg/handlers"
	clubhttp "github.com/gojou/goof/pkg/rest/club"
	fencerhttp "github.com/gojou/goof/pkg/rest/fencer"
	"github.com/gojou/goof/pkg/services/club"
	"github.com/gojou/goof/pkg/services/fencer"

	"github.com/gorilla/mux"
)

// Routing defines the API for goof
func Routing(r *mux.Router) {
	fs := fencer.NewService()
	cs := club.NewService()
	fhs := fencerhttp.NewServer(fs, r)
	chs := clubhttp.NewServer(cs, r)

	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/club/list", chs.HandleHTTP())
	r.HandleFunc("/club/jlist", chs.HandleJSON())
	r.HandleFunc("/fencer/list", fhs.HandleHTTP())
	r.HandleFunc("/fencer/jlist", fhs.HandleJSON())
	r.HandleFunc("/club/add", chs.HandleAdd())
	r.HandleFunc("/fencer/add", fhs.HandleAdd())
}
