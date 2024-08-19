package usecase

import (
	"backend/modules/inventory-operation"
	resultModels "backend/modules/inventory-operation/delivery/http/models"
	"backend/modules/inventory-operation/repository/mysql"
	"backend/modules/inventory-operation/repository/mysql/models"
	"backend/modules/user"
	users_model "backend/modules/user/repository/mysql/models"
	"gorm.io/gorm"
)

type InventoryOperationsModuleUC struct {
	mysqlRepo inventory_operation.MysqlRepoInterface
	UserUC    user.UseCaseInterface
}

func NewInventoryOperationsModuleUseCase(mysqlRepo *gorm.DB) *InventoryOperationsModuleUC {
	return &InventoryOperationsModuleUC{
		mysqlRepo: mysql.NewMysqlRepository(mysqlRepo),
	}
}

func (uc *InventoryOperationsModuleUC) GetTransactions(filter models.OperationFilter) (resultModels.OperationsResult, error) {
	return uc.mysqlRepo.GetTransactions(filter)
}

func (uc *InventoryOperationsModuleUC) GetInventory(filter models.InventoryFilters) (resultModels.InventoryOperationResultsNested, error) {
	return uc.mysqlRepo.GetInventory(filter)
}

func (uc *InventoryOperationsModuleUC) SetUserUseCase(userUC user.UseCaseInterface) {
	uc.UserUC = userUC
}

func (uc *InventoryOperationsModuleUC) BindUsersWithTransactions(usersFilter users_model.UsersObject, operations resultModels.OperationsResult) (resultModels.OperationsResult, error) {
	return uc.bindUsersWithTransactions(usersFilter, operations)
}

func (uc *InventoryOperationsModuleUC) BindUsersWithInventory(usersFilter users_model.UsersObject, inventory resultModels.InventoryOperationResultsNested) (resultModels.InventoryOperationResultsNested, error) {
	return uc.bindUsersWithInventory(usersFilter, inventory)
}
