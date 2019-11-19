package cmd

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"go-web/component/log"
	"go-web/errors"
	"go-web/route"
	"net/http"
	"runtime/debug"
	"time"
)

func Api(c *cli.Context) {
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

type ErrResponse struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Stack string      `json:"stack"`
}

func (e *ErrResponse) Error() string {
	return fmt.Sprintf("code: %d  errMsg: %s  stack: %s", e.Code, e.Msg, e.Stack)
}

//处理自定义返回错误
func httpErrorHandler(err error, c echo.Context) {

	var (
		code = http.StatusInternalServerError
		msg  *ErrResponse
	)

	var errFlag string
	//自定义错误
	if e, ok := err.(*errors.HttpError); ok {
		errFlag = " service error --> "
		code = e.HttpState
		msg = &ErrResponse{
			Code:  e.Code,
			Msg:   e.ErrMsg,
			Stack: e.Stack,
		}

	} else if e, ok := err.(*echo.HTTPError); ok {
		// echo 错误
		errFlag = " echo frame error --> "
		code = e.Code
		msg = &ErrResponse{
			Code:  e.Code,
			Msg:   e.Message,
			Stack: string(debug.Stack()),
		}
	} else {
		//其他错误，如运行时发生panic
		errFlag = " service panic --> "
		msg = &ErrResponse{
			Code:  code,
			Msg:   "panic",
			Stack: string(debug.Stack()),
		}

	}

	log.Sugar.Error(errFlag, msg)

	//线上环境不暴露堆栈信息
	if viper.Get("system.runtime") == "online" {
		msg.Stack = ""
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			log.Sugar.Error("Send response  err  ", err)
		}
	}
}
