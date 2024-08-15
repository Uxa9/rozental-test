package usecase

import (
	"backend/modules/inventory-operation"
	"backend/modules/inventory-operation/repository/mysql"
	"backend/modules/inventory-operation/repository/mysql/models"
	"gorm.io/gorm"
)

type InventoryOperationsModuleUC struct {
	mysqlRepo inventory_operation.MysqlRepoInterface
}

func NewInventoryOperationsModuleUseCase(mysqlRepo *gorm.DB) *InventoryOperationsModuleUC {
	return &InventoryOperationsModuleUC{
		mysqlRepo: mysql.NewMysqlRepository(mysqlRepo),
	}
}

func (uc *InventoryOperationsModuleUC) GetTransactions(filter models.OperationFilter) (models.OperationsResult, error) {
	return uc.mysqlRepo.GetTransactions(filter)
}

func (uc *InventoryOperationsModuleUC) GetInventory(filter models.InventoryFilters) (models.InventoryOperationResultsNested, error) {
	return uc.mysqlRepo.GetInventory(filter)
}
