package inventory_operation

import (
	resultModels "backend/modules/inventory-operation/delivery/http/models"
	"backend/modules/inventory-operation/repository/mysql/models"
)

type MysqlRepoInterface interface {
	GetTransactions(filter models.OperationFilter) (resultModels.OperationsResult, error)
	GetInventory(filter models.InventoryFilters) (resultModels.InventoryOperationResultsNested, error)
}
