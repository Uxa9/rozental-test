package app

import "gorm.io/gorm"

type App struct {
	Mysql *gorm.DB
}

func NewApp(gormMysql *gorm.DB) *App {
	return &App{
		Mysql: gormMysql,
	}
}

func GetApp(mysql *gorm.DB) *App {
	return NewApp(mysql)
}
