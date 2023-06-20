package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"openapi-gene-sample/controller"
	api "openapi-gene-sample/gen/go"
)

func main() {
	e := echo.New()
	e.Use(middleware.BodyDump(func(context echo.Context, req []byte, res []byte) {
		log.Info(string(req))
		log.Info(string(res))
	}))

	petController := &controller.PetController{}
	api.RegisterHandlers(e, petController)

	e.Logger.Fatal(e.Start(":8080"))
}
