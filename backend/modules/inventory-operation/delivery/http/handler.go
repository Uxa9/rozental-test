package http

import (
	inventory_operation "backend/modules/inventory-operation"
	resultModels "backend/modules/inventory-operation/delivery/http/models"
	"backend/modules/inventory-operation/repository/mysql/models"
	"backend/modules/user"
	users_model "backend/modules/user/repository/mysql/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	userUC               user.UseCaseInterface
	inventoryOperationUC inventory_operation.UseCaseInterface
}

func NewHandler(inventoryOperationUC inventory_operation.UseCaseInterface, userUC user.UseCaseInterface) *Handler {
	return &Handler{
		inventoryOperationUC: inventoryOperationUC,
		userUC:               userUC,
	}
}

func (handler *Handler) GetInventoryOperation(context *gin.Context) {
	result := resultModels.OperationsResult{}
	filter := models.OperationFilter{}
	usersFilters := users_model.UsersObject{}

	if err := context.ShouldBindJSON(&filter); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.inventoryOperationUC.GetTransactions(filter)

	for item := range result {
		usersFilters.Ids = append(usersFilters.Ids, result[item].DstExecutor.ExecutorId)
		usersFilters.Ids = append(usersFilters.Ids, result[item].SrcExecutor.ExecutorId)
	}

	nestedResult, _ := handler.inventoryOperationUC.BindUsersWithTransactions(usersFilters, result)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": &nestedResult})

	return
}

func (handler *Handler) GetInventory(context *gin.Context) {
	result := resultModels.InventoryOperationResultsNested{}
	filter := models.InventoryFilters{}
	usersFilters := users_model.UsersObject{}

	if err := context.ShouldBindJSON(&filter); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.inventoryOperationUC.GetInventory(filter)

	for item := range result {
		usersFilters.Ids = append(usersFilters.Ids, result[item].Executor.ExecutorId)

		for transaction := range result[item].Transactions {
			usersFilters.Ids = append(usersFilters.Ids, result[item].Transactions[transaction].DstExecutor.ExecutorId)
		}
	}

	nestedResult, _ := handler.inventoryOperationUC.BindUsersWithInventory(usersFilters, result)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": &nestedResult})

	return
}
