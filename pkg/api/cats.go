package api

import (
	"encoding/json"
	"errors"
	"example/pkg/models"
	memory_storage "example/pkg/storage/memory-storage"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func (h Handler) getCats(w http.ResponseWriter, r *http.Request) {
	cats, err := h.storage.GetAllCats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(cats); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h Handler) getCat(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	cat, err := h.storage.GetCatByName(name)
	if err != nil {
		if errors.Is(err, memory_storage.CatNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cat); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (h Handler) createCat(w http.ResponseWriter, r *http.Request) {
	var cat models.Cat
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	newName, err := modifyCatName(cat.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cat.Name = newName

	cat, err = h.storage.CreateCat(cat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(cat); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h Handler) removeCat(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	err := h.storage.RemoveCatByName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func modifyCatName(name string) (string, error) {
	if name == "" {
		return "", errors.New("cat name is empty")
	}
	return fmt.Sprintf("%s:%s", strings.ToUpper(name), time.Now().Weekday().String()), nil
}
