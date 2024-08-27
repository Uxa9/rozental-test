package usecase

import (
	resultModels "backend/modules/inventory-operation/delivery/http/models"
	"backend/modules/status/delivery/http/model"
)

func (uc *InventoryOperationsModuleUC) bindStatusWithTransactions(operations resultModels.OperationsResult) (resultModels.OperationsResult, error) {

	status, err := uc.StatusUC.GetStatus(model.GetStatusInput{})

	var statusMap = make(map[int]string)

	for _, stat := range status.List {
		statusMap[stat.Id] = stat.Name
	}

	for i := range operations {
		row := &operations[i]
		row.Status.StatusName = statusMap[row.Status.StatusId]
	}

	return operations, err
}

func (uc *InventoryOperationsModuleUC) bindStatusWithInventory(inventory resultModels.InventoryOperationResultsNested) (resultModels.InventoryOperationResultsNested, error) {

	status, err := uc.StatusUC.GetStatus(model.GetStatusInput{})

	var statusMap = make(map[int]string)

	for _, stat := range status.List {
		statusMap[stat.Id] = stat.Name
	}

	for i := range inventory {
		row := &inventory[i]

		for j := range row.Transactions {
			transaction := &row.Transactions[j]
			transaction.Status.StatusName = statusMap[transaction.Status.StatusId]
		}
	}

	return inventory, err
}
