package routers

import (
	"echo-clean/src/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	root := e.Group("api/v1")

	controllers.AddMemberRoutes(root)
}
