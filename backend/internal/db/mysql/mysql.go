package mysql

import (
	"backend/internal/config"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlService() *gorm.DB {

	mysqlDBConn := getMysqlConfigYaml(config.Yaml)

	var gormConfig gorm.Config

	gormMysql, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlDBConn,
	}), &gormConfig)

	if err != nil {
		panic(err)
	}

	return gormMysql
}

func getMysqlConfig(user, password, host, port, dbName string) (dbConn *sql.DB) {

	dbConn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user,
		password,
		host,
		port,
		dbName,
	))

	if err != nil {
		panic(err)
	}

	return
}

// overload 🥴🥴🥴
func getMysqlConfigYaml(sqlConfig config.MySqlConfig) (dbConn *sql.DB) {

	dbConn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		sqlConfig.User,
		sqlConfig.Password,
		sqlConfig.Host,
		sqlConfig.Port,
		sqlConfig.Db,
	))

	if err != nil {
		panic(err)
	}

	return
}
