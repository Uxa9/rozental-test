package mysql

import (
	"backend/internal/cmd/entity"
	"backend/modules/inventory-operation/repository/mysql/models"
)

func (r *MysqlRepository) GetInventory(filter models.InventoryFilters) (models.InventoryOperationResultsNested, error) {

	result := models.InventoryOperationResults{}
	client := r.client.Table(entity.Inventory.Table)

	client.Select("" +
		"`inventory`.`id`, " +
		"`inventory`.`name`, " +
		"`inventory`.`id_executor` AS `idSrcExecutor`, " +
		"u1.name AS `nameSrcExecutor`, " +
		"u2.name AS `nameDstExecutor`, " +
		"`inventory_operation_status`.`name` AS `statusName`, " +
		"`inventory_operations`.`id` AS `operationId`, " +
		"`inventory_operations`.`request_time`, " +
		"`inventory_operations`.`dst_executor` AS `idDstExecutor`, " +
		"`inventory_operation_status`.`id` AS `idStatus`, " +
		"`inventory_operation_status`.`name` AS `nameStatus`, " +
		"`inventory_operations`.`status_time` ",
	)

	client.Joins("" +
		"JOIN `user` u1 ON `u1`.`id` = `inventory`.`id_executor`" +
		"JOIN `inventory_operations_detail` ON `inventory_operations_detail`.`inventory_id` = `inventory`.`id`" +
		"JOIN `inventory_operations` ON `inventory_operations`.`id` = `inventory_operations_detail`.`operation_id`" +
		"JOIN `user` u2 ON `u2`.`id` = `inventory_operations`.`dst_executor`" +
		"JOIN `inventory_operation_status` ON `inventory_operation_status`.`id` = `inventory_operations`.`status`",
	)

	if len(filter.Id) == 0 && len(filter.Name) == 0 && len(filter.ExecutorId) == 0 {
		client.Find(&result)
	} else {
		if len(filter.Id) != 0 {
			client.Where(filter.Id)
		}

		if len(filter.Name) != 0 {
			client.Where("`inventory`.`name` IN (?)", filter.Name)
		}

		if len(filter.ExecutorId) != 0 {
			client.Where("`inventory_operations`.`dst_executor` IN (?)", filter.ExecutorId)
		}

		if len(filter.Status) != 0 {
			client.Where("`inventory_operation_status`.`name` IN (?)", filter.Status)
		}

		client.Find(&result)
	}

	var a []models.InventoryOperationResultNested

	for i, _ := range result {
		double := false
		for j, _ := range a {
			if a[j].Id == result[i].Id {
				double = true
				a[j].Transactions = append(a[j].Transactions, result[i].Transaction)
				break
			}
		}
		if !double {
			a = append(a, models.InventoryOperationResultNested{
				Id:           result[i].Id,
				Name:         result[i].Name,
				Executor:     result[i].Executor,
				Transactions: []models.InventoryOperation{result[i].Transaction},
			})
		}
	}

	return a, client.Error
}
