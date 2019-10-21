package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	appUser "go-web/api/app/account/api"
	webUser "go-web/api/web/account/api"
	custMidd "go-web/middleware"
)

func Route(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KB
		DisableStackAll:   false,
		DisablePrintStack: true,
	}))

	web := e.Group("webapi", custMidd.TokenMiddleware)
	web.GET("/user/manage", webUser.UserHandler.UserManage())

	app := e.Group("api")
	app.POST("/user", appUser.UserHandler.CreateUser())
}
