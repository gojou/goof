package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gojou/goof/cmd/routes"
	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}
func run() (e error) {
	r := mux.NewRouter()

	routes.Routing(r)

	// Critical to work on AppEngine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	return (http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
