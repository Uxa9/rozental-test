package usecase

import (
	resultModels "backend/modules/inventory-operation/delivery/http/models"
	users_model "backend/modules/user/repository/mysql/models"
)

func (uc *InventoryOperationsModuleUC) bindUsersWithTransactions(usersFilter users_model.UsersObject, operations resultModels.OperationsResult) (resultModels.OperationsResult, error) {

	users, err := uc.UserUC.GetUser(usersFilter)

	var userMap = make(map[int]string)

	for _, user := range users {
		userMap[user.Id] = user.Name
	}

	for i := range operations {
		row := &operations[i]
		row.SrcExecutor.ExecutorName = userMap[row.SrcExecutor.ExecutorId]
		row.DstExecutor.ExecutorName = userMap[row.DstExecutor.ExecutorId]
	}

	return operations, err
}

func (uc *InventoryOperationsModuleUC) bindUsersWithInventory(usersFilter users_model.UsersObject, inventory resultModels.InventoryOperationResultsNested) (resultModels.InventoryOperationResultsNested, error) {

	users, err := uc.UserUC.GetUser(usersFilter)

	var userMap = make(map[int]string)

	for _, user := range users {
		userMap[user.Id] = user.Name
	}

	for i := range inventory {
		row := &inventory[i]
		row.Executor.ExecutorName = userMap[row.Executor.ExecutorId]

		for j := range row.Transactions {
			transaction := &row.Transactions[j]
			transaction.DstExecutor.ExecutorName = userMap[transaction.DstExecutor.ExecutorId]
		}
	}

	return inventory, err
}
