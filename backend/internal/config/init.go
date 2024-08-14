package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Init() error {
	viper.AddConfigPath("./internal/config")

	_, err := os.Stat("./internal/config/config.yml")

	if err == nil {
		//viper.SetConfigName("config")
		viper.SetConfigFile("./internal/config/config.yml")
	} else {
		panic("config file is not exist")
	}
	fmt.Println("read config file:", viper.ConfigFileUsed())
	return viper.ReadInConfig()
}
