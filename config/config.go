package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Application struct {
		Name string `mapstructure:"name"`
		Port string `mapstructure:"port"`
	}
	Jwt struct {
		Expire int64  `mapstructure:"expire"`
		Secret string `mapstructure:"secret"`
	}
	Aes struct {
		Secret string `mapstructure:"secret"`
	}
	Mysql struct {
		Username  string `mapstructure:"username"`
		Password  string `mapstructure:"password"`
		Addr      string `mapstructure:"addr"`
		Port      string `mapstructure:"port"`
		Databases string `mapstructure:"databases"`
		Charset   string `mapstructure:"charset"`
	}
	Huawei struct {
		Oauth2URL    string `mapstructure:"oauth2_url"`
		ClientID     string `mapstructure:"client_id"`
		ClientSecret string `mapstructure:"client_secret"`
		AccountUrl   string `mapstructure:"account_url"`
	}
}

var GlobalConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: , %v", err)
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("解析配置失败: , %v", err)
	}
}
