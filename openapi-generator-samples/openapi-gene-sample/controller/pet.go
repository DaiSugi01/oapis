package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"

	gen "oapi-sample/gen"
)

type PetController struct{}

func (p *PetController) FindPetById(c echo.Context) error {
	tag := "dog"
	pet := gen.Pet{
		Id:   1,
		Name: "Max",
		Tag:  &tag,
	}
	return c.JSON(http.StatusOK, pet)
}
