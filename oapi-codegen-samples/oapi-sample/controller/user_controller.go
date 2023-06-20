package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"oapi-sample/api"
)

var users = map[int]api.User{
	1: {Id: 1, Name: "Daiki"},
	2: {Id: 2, Name: "Sugihara"},
}

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) FindUsers(ctx echo.Context) error {
	user, ok := users[1]
	if !ok {
		return nil
	}

	resp, err := json.Marshal(user)
	if err != nil {
		return nil
	}

	return ctx.JSON(http.StatusOK, resp)
}
