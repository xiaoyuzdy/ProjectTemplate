package models

import (
	"go-web/component/db"
	"go-web/utils"
)

var UserHandler = User{}

type User struct {
	Id          int64          `gorm:"column:id" json:"id"`
	Account     string         `gorm:"column:account;type:varchar(55);comment:'账户'" json:"account"`
	AccountType uint8          `gorm:"column:account_type;type:tinyint unsigned;default:1;comment:'1--> 普通用户，2-->会员'" json:"account_type"`
	UserName    string         `gorm:"column:user_name;type:varchar(55);comment:'用户昵称'" json:"user_name"`
	CreatedAt   utils.JSONTime `gorm:"column:created_at" json:"-"`
	UpdatedAt   utils.JSONTime `gorm:"column:updated_at" json:"-"`
	DeletedAt   utils.JSONTime `gorm:"column:deleted_at" json:"-"`
}

func (u *User) TableName() string {
	return "user"
}

func (*User) PluckColumnByWhere(column string, value interface{}, query interface{}, args []interface{}) error {
	return db.Orm.Model(User{}).Where(query, args...).Pluck(column, value).Error
}

func (*User) QueryLastByWhere(user *User, query interface{}, args []interface{}) error {
	return db.Orm.Where(query, args...).Last(user).Error
}

func (*User) QueryBatchByWhere(query interface{}, args []interface{}) ([]User, error) {
	var list []User
	return list, db.Orm.Where(query, args...).Find(list).Error
}
