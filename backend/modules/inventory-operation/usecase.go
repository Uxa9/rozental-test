package inventory_operation

import (
	resultModels "backend/modules/inventory-operation/delivery/http/models"
	"backend/modules/inventory-operation/repository/mysql/models"
	users_model "backend/modules/user/repository/mysql/models"
)

type UseCaseInterface interface {
	GetTransactions(filter models.OperationFilter) (resultModels.OperationsResult, error)
	GetInventory(filter models.InventoryFilters) (resultModels.InventoryOperationResultsNested, error)
	BindUsersWithTransactions(usersFilter users_model.UsersObject, operations resultModels.OperationsResult) (resultModels.OperationsResult, error)
	BindUsersWithInventory(usersFilter users_model.UsersObject, inventory resultModels.InventoryOperationResultsNested) (resultModels.InventoryOperationResultsNested, error)
	BindStatusWithTransactions(operations resultModels.OperationsResult) (resultModels.OperationsResult, error)
	BindStatusWithInventory(inventory resultModels.InventoryOperationResultsNested) (resultModels.InventoryOperationResultsNested, error)
}
