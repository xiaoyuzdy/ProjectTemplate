package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Route(e *echo.Echo) {
	e.GET("/", func(context echo.Context) error {
		return context.JSON(200, "service run success")
	})

	e.Use(middleware.CORS())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KB
		DisableStackAll:   false,
		DisablePrintStack: true,
	}))
	//web := e.Group("webapi", custMidd.TokenMiddleware)
	//web.GET("/user/manage", webUser.UserHandler.UserManage())

	//app := e.Group("api")
	//app.POST("/user", appUser.UserHandler.CreateUser())
}
