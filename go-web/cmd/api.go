package cmd

import (
	"github.com/labstack/echo"
	"github.com/urfave/cli"
	"go-web/route"
	"net/http"
	"time"
)

func Api(*cli.Context) {
	e := echo.New()
	e.HTTPErrorHandler = httpErrorHandler
	route.Route(e)
	s := &http.Server{
		Addr:         ":8099",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}

//处理自定义返回错误
func httpErrorHandler(err error, c echo.Context) {


}
