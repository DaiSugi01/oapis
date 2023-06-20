package controller

import (
	"github.com/labstack/echo/v4"
	"oapi-sample/api"
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

func (c *MainController) FindPet(ctx echo.Context) error {
	return c.PetController.FindPet(ctx)
}

func (c *MainController) FindPetWithQueryParam(ctx echo.Context, params api.FindPetWithQueryParamParams) error {
	return c.PetController.FindPetWithQueryParam(ctx, params)
}

func (c MainController) FindPetById(ctx echo.Context, id int64) error {
	return c.PetController.FindPetById(ctx, id)
}

func (c *MainController) FindUsers(ctx echo.Context) error {
	return c.UserController.FindUsers(ctx)
}
