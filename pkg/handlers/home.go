package handlers

import (
	"fmt"
	"net/http"
)

//Home presents a very basic home page.
func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello %s\n", "world")

}
