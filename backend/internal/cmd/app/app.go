package app

import (
	"backend/modules/user"
	userModuleUseCase "backend/modules/user/usecase"
	"gorm.io/gorm"
)

type App struct {
	Mysql        *gorm.DB
	UserModuleUC user.UseCaseInterface
}

func NewApp(gormMysql *gorm.DB) *App {
	return &App{
		Mysql:        gormMysql,
		UserModuleUC: userModuleUseCase.NewUserModuleUseCase(gormMysql),
	}
}

func GetApp(mysql *gorm.DB) *App {
	return NewApp(mysql)
}
