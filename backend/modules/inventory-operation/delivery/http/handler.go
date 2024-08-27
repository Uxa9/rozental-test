package http

import (
	inventory_operation "backend/modules/inventory-operation"
	resultModels "backend/modules/inventory-operation/delivery/http/models"
	"backend/modules/inventory-operation/repository/mysql/models"
	"backend/modules/status"
	"backend/modules/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	userUC               user.UseCaseInterface
	inventoryOperationUC inventory_operation.UseCaseInterface
	statusUC             status.UseCaseInterface
}

func NewHandler(
	inventoryOperationUC inventory_operation.UseCaseInterface,
	userUC user.UseCaseInterface,
	statusUC status.UseCaseInterface,
) *Handler {
	return &Handler{
		inventoryOperationUC: inventoryOperationUC,
		userUC:               userUC,
		statusUC:             statusUC,
	}
}

func (handler *Handler) GetInventoryOperation(context *gin.Context) {
	result := resultModels.OperationsResult{}
	filter := models.OperationFilter{}

	if err := context.ShouldBindJSON(&filter); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.inventoryOperationUC.GetTransactions(filter)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": &result})

	return
}

func (handler *Handler) GetInventory(context *gin.Context) {
	result := resultModels.InventoryOperationResultsNested{}
	filter := models.InventoryFilters{}

	if err := context.ShouldBindJSON(&filter); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.inventoryOperationUC.GetInventory(filter)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": &result})

	return
}
