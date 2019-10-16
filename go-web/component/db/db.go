package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"go-web/component/log"
	"go-web/utils"
	"gopkg.in/redis.v5"
	"time"
)

var Redis *redis.Client
var Orm *gorm.DB

func InitRedis() {
	host := viper.GetString("redisServer.host")
	post := viper.GetString("redisServer.port")
	auth := viper.GetString("redisServer.auth")
	encryption := viper.GetInt("redisServer.encryption")
	dbIndex := viper.GetInt("redisDB.demo")
	if encryption == 1 {
		auth = utils.GetMD5(auth)
	}
	for {
		options := redis.Options{Addr: host + post, Password: auth, DB: dbIndex, MaxRetries: 3}
		client := redis.NewClient(&options)
		_, err := client.Ping().Result()
		if err == nil {
			Redis = client
			break
		}
		log.Sugar.Error("redis connection exception! 5 seconds to retry, err = %v \r\n", err)
		time.Sleep(time.Second * 5)
	}
}

func InitMysql() {
	dbname := viper.GetString("mysqlServer.dbname")
	host := viper.GetString("mysqlServer.host")
	port := viper.GetString("mysqlServer.port")
	user := viper.GetString("mysqlServer.username")
	pwd := viper.GetString("mysqlServer.password")
	gormLog := viper.GetBool("log.gorm_log")
	for {
		orm, err := gorm.Open("mysql", user+":"+pwd+"@("+host+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
		if err == nil {
			if !gormLog {
				orm.LogMode(gormLog)
			}
			orm.DB().SetMaxOpenConns(100)
			orm.DB().SetMaxIdleConns(50)
			Orm = orm
			break
		}
		log.Sugar.Error("Database connection exception! 5 seconds to retry, err = %v \r\n", err)
		time.Sleep(5 * time.Second)
	}
}
