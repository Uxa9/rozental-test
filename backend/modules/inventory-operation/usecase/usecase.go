package usecase

import (
	"backend/modules/inventory-operation"
	resultModels "backend/modules/inventory-operation/delivery/http/models"
	"backend/modules/inventory-operation/repository/mysql"
	"backend/modules/inventory-operation/repository/mysql/models"
	"backend/modules/status"
	"backend/modules/user"
	users_model "backend/modules/user/repository/mysql/models"
	"gorm.io/gorm"
)

type InventoryOperationsModuleUC struct {
	mysqlRepo inventory_operation.MysqlRepoInterface
	UserUC    user.UseCaseInterface
	StatusUC  status.UseCaseInterface
}

func NewInventoryOperationsModuleUseCase(mysqlRepo *gorm.DB) *InventoryOperationsModuleUC {
	return &InventoryOperationsModuleUC{
		mysqlRepo: mysql.NewMysqlRepository(mysqlRepo),
	}
}

func (uc *InventoryOperationsModuleUC) GetTransactions(filter models.OperationFilter) (resultModels.OperationsResult, error) {
	usersFilters := users_model.UsersObject{}
	result, err := uc.mysqlRepo.GetTransactions(filter)

	for item := range result {
		usersFilters.Ids = append(usersFilters.Ids, result[item].DstExecutor.ExecutorId)
		usersFilters.Ids = append(usersFilters.Ids, result[item].SrcExecutor.ExecutorId)
	}

	nestedResult, err := uc.BindUsersWithTransactions(usersFilters, result)

	if err != nil {
		return resultModels.OperationsResult{}, err
	}

	nestedResult, err = uc.BindStatusWithTransactions(result)

	if err != nil {
		return resultModels.OperationsResult{}, err
	}

	return nestedResult, nil
}

func (uc *InventoryOperationsModuleUC) GetInventory(filter models.InventoryFilters) (resultModels.InventoryOperationResultsNested, error) {
	usersFilters := users_model.UsersObject{}
	result, err := uc.mysqlRepo.GetInventory(filter)

	if err != nil {
		return resultModels.InventoryOperationResultsNested{}, err
	}

	for item := range result {
		usersFilters.Ids = append(usersFilters.Ids, result[item].Executor.ExecutorId)

		for transaction := range result[item].Transactions {
			usersFilters.Ids = append(usersFilters.Ids, result[item].Transactions[transaction].DstExecutor.ExecutorId)
		}
	}

	nestedResult, err := uc.BindUsersWithInventory(usersFilters, result)

	if err != nil {
		return resultModels.InventoryOperationResultsNested{}, err
	}

	nestedResult, err = uc.BindStatusWithInventory(result)

	if err != nil {
		return resultModels.InventoryOperationResultsNested{}, err
	}

	return nestedResult, nil
}

func (uc *InventoryOperationsModuleUC) SetUserUseCase(userUC user.UseCaseInterface) {
	uc.UserUC = userUC
}

func (uc *InventoryOperationsModuleUC) SetStatusUseCase(statusUC status.UseCaseInterface) {
	uc.StatusUC = statusUC
}

func (uc *InventoryOperationsModuleUC) BindUsersWithTransactions(usersFilter users_model.UsersObject, operations resultModels.OperationsResult) (resultModels.OperationsResult, error) {
	return uc.bindUsersWithTransactions(usersFilter, operations)
}

func (uc *InventoryOperationsModuleUC) BindUsersWithInventory(usersFilter users_model.UsersObject, inventory resultModels.InventoryOperationResultsNested) (resultModels.InventoryOperationResultsNested, error) {
	return uc.bindUsersWithInventory(usersFilter, inventory)
}

func (uc *InventoryOperationsModuleUC) BindStatusWithTransactions(operations resultModels.OperationsResult) (resultModels.OperationsResult, error) {
	return uc.bindStatusWithTransactions(operations)
}

func (uc *InventoryOperationsModuleUC) BindStatusWithInventory(inventory resultModels.InventoryOperationResultsNested) (resultModels.InventoryOperationResultsNested, error) {
	return uc.bindStatusWithInventory(inventory)
}
