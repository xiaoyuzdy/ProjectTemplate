package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

func InitViper() {
	config(viper.GetViper())
	filename := `config.toml`
	_, err := os.Stat(filename)
	if (err == nil) || (os.IsExist(err)) {
		viper.SetConfigFile(filename)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		viper.WatchConfig()
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func config(viper *viper.Viper) {
	//设置默认值
	//file
	viper.SetDefault("staticFile.ip", "http://127.0.0.1")
	viper.SetDefault("staticFile.port", ":8081")
	//mysql
	viper.SetDefault("mysqlServer.host", "127.0.0.1")
	viper.SetDefault("mysqlServer.port", ":3306")
	viper.SetDefault("mysqlServer.dbname", "default")
	viper.SetDefault("mysqlServer.username", "default")
	viper.SetDefault("mysqlServer.password", "default")
	//redis
	viper.SetDefault("redisServer.host", "127.0.0.1")
	viper.SetDefault("redisServer.port", ":6379")
	viper.SetDefault("redisServer.auth", "")
	viper.SetDefault("redisServer.encryption", 0)
	viper.SetDefault("redisDB.demo", 0)
}
