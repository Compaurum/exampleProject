package memory_storage

import (
	"errors"
	"example/pkg/models"
)

var CatNotFound = errors.New("cat not found")

type MemoryStorage map[string]models.Cat

func (s MemoryStorage) fefe(name string) (models.Cat, error) {
	return s[name], nil
}

func (s MemoryStorage) GetCatByName(name string) (models.Cat, error) {
	cat, ok := s[name]
	if !ok {
		return cat, CatNotFound
	}
	return cat, nil
}

func (s MemoryStorage) GetAllCats() ([]models.Cat, error) {
	cats := make([]models.Cat, 0, len(s))
	for _, cat := range s {
		cats = append(cats, cat)
	}

	return cats, nil
}

func (s MemoryStorage) RemoveCatByName(name string) error {
	delete(s, name)
	return nil
}

func (s MemoryStorage) CreateCat(cat models.Cat) (models.Cat, error) {
	s[cat.Name] = cat

	return cat, nil
}
