package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"oapi-sample/api"
	"oapi-sample/controller"
)

func main() {
	e := echo.New()
	e.Use(middleware.BodyDump(func(context echo.Context, req []byte, res []byte) {
		log.Info(string(req))
		log.Info(string(res))
	}))

	mainController := controller.NewMainController(
		controller.NewPetController(),
		controller.NewUserController(),
	)

	api.RegisterHandlers(e, mainController)

	e.Logger.Fatal(e.Start(":8080"))
}
