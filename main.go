package main

import (
	"echo-clean/src/configs"
	"echo-clean/src/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	configs.CreateDatabaseConnection()
	routers.SetupRouter(e)

	e.Logger.Fatal(e.Start(":" + configs.GetEnv("PORT")))
}
