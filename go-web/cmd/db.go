package cmd

import (
	"go-web/component/db"
	"go-web/component/log"
	"go-web/models"
)

func CreateTable() {
	log.Sugar.Debug("开始初始化数据库")

	db.Orm.Set("gorm:table_options", "CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&models.User{})

	log.Sugar.Debug("数据库初始化完成")
}
