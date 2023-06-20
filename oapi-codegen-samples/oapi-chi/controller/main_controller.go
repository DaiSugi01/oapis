package controller

import (
	"net/http"
	"oapi-chi/api"
)

type MainController struct {
	PetController  *PetController
	UserController *UserController
}

func NewMainController(petController *PetController, userController *UserController) *MainController {
	return &MainController{
		petController,
		userController,
	}
}

func (c *MainController) FindPet(w http.ResponseWriter, r *http.Request) {
	c.PetController.FindPet(w, r)
}

func (c *MainController) FindPetWithQueryParam(w http.ResponseWriter, r *http.Request, params api.FindPetWithQueryParamParams) {
	c.PetController.FindPetWithQueryParam(w, r, params)
}

func (c *MainController) FindPetById(w http.ResponseWriter, r *http.Request, id int64) {
	c.PetController.FindPetById(w, r, id)
}

func (c *MainController) FindUsers(w http.ResponseWriter, r *http.Request) {
	c.UserController.FindUsers(w, r)
}
