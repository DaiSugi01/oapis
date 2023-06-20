package controller

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"oapi-sample/api"
	chi_api "oapi-sample/api/chi"
)

var pets = map[int]api.Pet{
	1: {Id: 1, Name: "Dog"},
	2: {Id: 2, Name: "Cat"},
}

type PetController struct{}

func NewPetController() *PetController {
	return &PetController{}
}

func (c *PetController) FindPet(ctx echo.Context) error {
	//pet, ok := pets[1]
	//if !ok {
	//	return nil
	//}
	client, err := chi_api.NewClientWithResponses("http://localhost:9090",
		chi_api.WithHTTPClient(http.DefaultClient),
		chi_api.WithRequestEditorFn(
			func(ctx context.Context, req *http.Request) error {
				// Here you can modify the request, for example, add headers, change URL, etc.
				return nil
			},
		),
	)
	if err != nil {
		return err
	}

	resp, err := client.FindPetWithResponse(ctx.Request().Context())
	if err != nil {
		return err
	}

	// Check the HTTP status code
	if resp.StatusCode() != http.StatusOK {
		return echo.NewHTTPError(resp.StatusCode(), fmt.Errorf("%v", "errror happened!"))
	}

	// Decode the response body into the `Pet` type
	pet := resp.JSON200
	return ctx.JSON(http.StatusOK, &pet)
}

func (c *PetController) FindPetWithQueryParam(ctx echo.Context, params api.FindPetWithQueryParamParams) error {
	return ctx.String(http.StatusOK, ctx.QueryParam("id"))
}

func (c PetController) FindPetById(ctx echo.Context, id int64) error {
	return ctx.JSON(http.StatusOK, id)
}
