package controller

import (
	"encoding/json"
	"net/http"
	"oapi-chi/api"
)

var pets = map[int]api.Pet{
	1: {Id: 1, Name: "Dog"},
	2: {Id: 2, Name: "Cat"},
}

type PetController struct{}

func NewPetController() *PetController {
	return &PetController{}
}

func (c *PetController) FindPet(w http.ResponseWriter, r *http.Request) {
	pet, ok := pets[int(1)]
	if !ok {
		http.Error(w, "Pet not found", http.StatusNotFound)
		return
	}
	resp, err := json.Marshal(pet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (c PetController) FindPetWithQueryParam(w http.ResponseWriter, r *http.Request, params api.FindPetWithQueryParamParams) {
	resp, err := json.Marshal(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (c *PetController) FindPetById(w http.ResponseWriter, r *http.Request, id int64) {
	resp, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
