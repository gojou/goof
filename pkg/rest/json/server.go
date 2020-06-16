package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ServeJSON serves the list of fencers in JSON format
func (s *service) ServeJSON(w http.ResponseWriter, r *http.Request) {
	flist, _ := s.r.ListFencers()
	json, _ := json.MarshalIndent(flist, "", "  ")
	fmt.Fprintln(w, string(json))

}
