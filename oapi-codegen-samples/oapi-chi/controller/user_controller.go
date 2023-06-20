package controller

import (
	"encoding/json"
	"net/http"
	"oapi-chi/api"
)

var users = map[int]api.User{
	1: {Id: 1, Name: "Daiki"},
	2: {Id: 2, Name: "Sugihara"},
}

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) FindUsers(w http.ResponseWriter, r *http.Request) {
	user, ok := users[1]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
