package club

import "github.com/gojou/goof/pkg/models"

//Storage provides the abstract structure for storing club info
type Storage interface {
	GetClub(id string) (*models.Club, error)
	AddClub(models.Club) (*models.Club, error)
	ListClubs() (*[]models.Club, error)
}
