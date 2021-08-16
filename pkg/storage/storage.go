package storage

import "example/pkg/models"

type Storage interface {
	GetCatByName(name string) (models.Cat, error)
	GetAllCats() ([]models.Cat, error)
	RemoveCatByName(name string) error
	CreateCat(cat models.Cat) (models.Cat, error)
}
