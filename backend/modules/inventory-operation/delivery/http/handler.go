package http

import (
	inventory_operation "backend/modules/inventory-operation"
	"backend/modules/inventory-operation/repository/mysql/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	inventoryOperationUC inventory_operation.UseCaseInterface
}

func NewHandler(inventoryOperationUC inventory_operation.UseCaseInterface) *Handler {
	return &Handler{
		inventoryOperationUC: inventoryOperationUC,
	}
}

func (handler *Handler) GetInventoryOperation(context *gin.Context) {
	result := models.OperationsResult{}
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

	context.JSON(http.StatusOK, gin.H{"data": result})

	return
}

func (handler *Handler) GetInventory(context *gin.Context) {
	result := models.InventoryOperationResultsNested{}
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

	context.JSON(http.StatusOK, gin.H{"data": result})

	return
}
