package main

import (
	"net/http"
	"oapi-chi/api"
	"oapi-chi/controller"
)

func main() {
	mainController := controller.NewMainController(
		controller.NewPetController(),
		controller.NewUserController(),
	)

	h := api.Handler(mainController)
	http.ListenAndServe(":9090", h)

	// net/httpを使用する場合
	//http.HandleFunc("/pets", mainController.FindPet)
	//http.HandleFunc("/users", mainController.FindUsers)
	//http.ListenAndServe(":8080", nil)
}
