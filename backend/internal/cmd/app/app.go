package app

import (
	"backend/modules/inventory-operation"
	inventoryOperationUseCase "backend/modules/inventory-operation/usecase"
	"backend/modules/status"
	statusModuleUseCase "backend/modules/status/usecase"
	"backend/modules/user"
	userModuleUseCase "backend/modules/user/usecase"
	"gorm.io/gorm"
)

type App struct {
	Mysql                *gorm.DB
	UserModuleUC         user.UseCaseInterface
	InventoryOperationUC inventory_operation.UseCaseInterface
	StatusUC             status.UseCaseInterface
}

func NewApp(gormMysql *gorm.DB) *App {
	inventoryOperationUC := inventoryOperationUseCase.NewInventoryOperationsModuleUseCase(gormMysql)
	userModuleUC := userModuleUseCase.NewUserModuleUseCase(gormMysql)
	statusModuleUC := statusModuleUseCase.NewStatusUseCase(gormMysql)

	inventoryOperationUC.SetUserUseCase(userModuleUC)
	inventoryOperationUC.SetStatusUseCase(statusModuleUC)

	return &App{
		Mysql:                gormMysql,
		UserModuleUC:         userModuleUC,
		InventoryOperationUC: inventoryOperationUC,
		StatusUC:             statusModuleUC,
	}
}

func GetApp(mysql *gorm.DB) *App {
	return NewApp(mysql)
}
