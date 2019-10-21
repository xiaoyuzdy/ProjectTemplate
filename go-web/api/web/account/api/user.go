package api

import "github.com/labstack/echo"

var UserHandler = &User{}

type User struct {
}

func (*User) UserManage() echo.HandlerFunc {
	return func(c echo.Context) error {

		return nil
	}
}
