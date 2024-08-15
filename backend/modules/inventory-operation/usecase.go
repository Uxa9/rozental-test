package inventory_operation

import "backend/modules/inventory-operation/repository/mysql/models"

type UseCaseInterface interface {
	GetTransactions(filter models.OperationFilter) (models.OperationsResult, error)
	GetInventory(filter models.InventoryFilters) (models.InventoryOperationResultsNested, error)
}
