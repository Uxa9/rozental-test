package main

import (
	"backend/internal/cmd/app"
	"backend/internal/cmd/start"
	"backend/internal/config"
	"backend/internal/db/mysql"
	"context"
	"gorm.io/gorm"
)

func initialization() (ctx context.Context, mysqlDb *gorm.DB) {
	config.FillConfig()

	ctx = context.TODO()
	mysqlDb = mysql.InitMysqlService()

	return
}

func main() {

	_, mysqlDb := initialization()

	application := app.GetApp(mysqlDb)

	start.StartServer(application)
}
