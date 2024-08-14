package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	MySqlConfig
	ServerConfig
}

type MySqlConfig struct {
	Host                  string
	Port                  string
	User                  string
	Password              string
	Db                    string
	MaxLengthArrayInQuery int
}

type ServerConfig struct {
	Port string
}

var Yaml Config

func FillConfig() {

	if err := Init(); err != nil {
		log.Fatalf(err.Error())
	}

	Yaml.MySqlConfig.Host = checkStringParameter("mysql.host")
	Yaml.MySqlConfig.Port = checkStringParameter("mysql.port")
	Yaml.MySqlConfig.User = checkStringParameter("mysql.user")
	Yaml.MySqlConfig.Password = checkStringParameter("mysql.password")
	Yaml.MySqlConfig.Db = checkStringParameter("mysql.db")

	Yaml.ServerConfig.Port = checkStringParameter("server.port")

	if viper.IsSet("MaxLengthArrayInQuery") {
		Yaml.MaxLengthArrayInQuery = viper.GetInt("MaxLengthArrayInQuery")
	} else {
		Yaml.MaxLengthArrayInQuery = 10000
	}
}

func checkStringParameter(parameter string) string {
	if !viper.IsSet(parameter) {
		fmt.Printf("%s is required\n", parameter)
		os.Exit(1)
	}
	return viper.GetString(parameter)
}

// overload 🥴🥴🥴
func checkIntParameter(parameter string) int {
	if !viper.IsSet(parameter) {
		fmt.Printf("%s is required\n", parameter)
		os.Exit(1)
	}
	return viper.GetInt(parameter)
}
