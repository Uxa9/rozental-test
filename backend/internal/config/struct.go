package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type MySqlConfig struct {
	Host                  string
	Port                  string
	User                  string
	Password              string
	Db                    string
	MaxLengthArrayInQuery int
}

var Yaml MySqlConfig

func FillConfig() {

	if err := Init(); err != nil {
		log.Fatalf(err.Error())
	}

	Yaml.Host = checkStringParameter("mysql.host")
	Yaml.Port = checkStringParameter("mysql.port")
	Yaml.User = checkStringParameter("mysql.user")
	Yaml.Password = checkStringParameter("mysql.password")
	Yaml.Db = checkStringParameter("mysql.db")

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
