package api

import (
	"github.com/labstack/echo"
	"go-web/errors"
	"go-web/models"
)

var UserHandler = &User{}

type User struct {
}

func (*User) UserManage() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := models.UserHandler.QueryLastByWhere(&models.User{}, "id = ?", []interface{}{100})
		if err != nil {
			return errors.ErrValidation(err)
		}
		return nil
	}
}
