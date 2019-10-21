package domain

import (
	"github.com/jinzhu/gorm"
	"go-web/errors"
	"go-web/models"
)

var UserHandler = &User{}

type User struct {
}

func (*User) CreateUser(account, pwd string) (string, error) {
	record := &models.User{}
	err := models.UserHandler.QueryLastByWhere(record, "account = ?", []interface{}{account})
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", errors.ErrInternalServerError(err)
	}
	if record.Id > 0 {
		return "", errors.ErrPrompt(errors.Code1000)
	}
	err = models.UserHandler.Insert(&models.User{
		Account:  account,
		Password: pwd,
	})
	if err != nil {
		return "", errors.ErrInternalServerError(err)
	}
	return "success", nil
}
