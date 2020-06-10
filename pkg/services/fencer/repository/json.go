package repository

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/gojou/goof/pkg/models"
)

//GetFencer returns a JSON representation of a Fencer
func GetFencer(id string) (*models.Fencer, error) {
	return nil, nil
}

//AddFencer saves a JSON representation of a Fencer
func AddFencer(fencer *models.Fencer) (*models.Fencer, error) {
	return nil, nil
}

//ListFencers returns a JSON representation of a all Fencers
func ListFencers() (*[]models.Fencer, error) {
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(fencers)
	s := buf.String()
	log.Println("JSON string: ", s)
	return &fencers, nil
}
