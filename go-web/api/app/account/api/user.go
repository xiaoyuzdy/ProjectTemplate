package api

import (
	"github.com/labstack/echo"
	"github.com/tidwall/gjson"
	"go-web/api/app/account/domain"
	"go-web/errors"
	"io/ioutil"
	"net/http"
)

var UserHandler = &User{}

type User struct {
}

func (*User) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		//TODO c.request   c.response  封装  参数校验
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return errors.ErrValidation(err)
		}
		json := gjson.Parse(string(body))
		account := json.Get("account").String()
		psw := json.Get("password").String()
		data, err := domain.UserHandler.CreateUser(account, psw)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"data": data,
			"msg":  "success",
		})
	}
}
