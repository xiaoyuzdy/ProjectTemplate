package component

import (
	"go-web/component/config"
	"go-web/component/db"
	"go-web/component/log"
)

func InitComponent() {
	initBase()
	db.InitMysql()
	db.InitRedis()
}

func initBase() {
	config.InitViper()
	log.InitLogs()
}
