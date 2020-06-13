package club

import (
	"net/http"

	"github.com/gojou/goof/pkg/models"
)

// Service todo
type Service interface {
	GetClub(id string) (*models.Club, error)
	AddClub(fencer *models.Club) (*models.Club, error)
	ListClubs() (*[]models.Club, error)
	Serve(w http.ResponseWriter, r *http.Request)
	ServeJSON(w http.ResponseWriter, r *http.Request)
}
