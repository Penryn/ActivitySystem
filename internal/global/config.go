package global

import (
	"log"

	"github.com/spf13/viper"
)

var Config = viper.New()

func init() {
	Config.AddConfigPath("conf")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.WatchConfig() // 监控配置文件变化
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

}