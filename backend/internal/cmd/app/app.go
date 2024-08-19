package app

import (
	"backend/modules/inventory-operation"
	inventoryOperationUseCase "backend/modules/inventory-operation/usecase"
	"backend/modules/user"
	userModuleUseCase "backend/modules/user/usecase"
	"gorm.io/gorm"
)

type App struct {
	Mysql                *gorm.DB
	UserModuleUC         user.UseCaseInterface
	InventoryOperationUC inventory_operation.UseCaseInterface
}

func NewApp(gormMysql *gorm.DB) *App {
	inventoryOperationUC := inventoryOperationUseCase.NewInventoryOperationsModuleUseCase(gormMysql)
	userModuleUC := userModuleUseCase.NewUserModuleUseCase(gormMysql)

	inventoryOperationUC.SetUserUseCase(userModuleUC)

	return &App{
		Mysql:                gormMysql,
		UserModuleUC:         userModuleUC,
		InventoryOperationUC: inventoryOperationUC,
	}
}

func GetApp(mysql *gorm.DB) *App {
	return NewApp(mysql)
}
