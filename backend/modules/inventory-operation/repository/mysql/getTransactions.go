package mysql

import (
	"backend/internal/cmd/entity"
	"backend/modules/inventory-operation/repository/mysql/models"
	"fmt"
	"time"
)

func (r *MysqlRepository) GetTransactions(filter models.OperationFilter) (result models.OperationsResult, err error) {

	result = models.OperationsResult{}
	client := r.client.Table(entity.InventoryOperation.Table)

	client.Select("" +
		"`inventory_operations`.`id` AS id, " +
		"`inventory_operations`.`src_executor` AS idSrcExecutor, " +
		"`inventory_operations`.`dst_executor` AS idDstExecutor, " +
		"`inventory_operations`.`request_time`, " +
		"`inventory_operations`.`status_time`, " +
		"`inventory_operations`.`status` AS statusId, " +
		"`inventory_operation_status`.`id` AS idStatus, " +
		"`inventory_operation_status`.`name` AS nameStatus, " +
		"`u1`.`name` AS nameSrcExecutor, " +
		"`u2`.`name` AS nameDstExecutor, " +
		"`inventory_operations_detail`.`inventory_id` AS idInventory, " +
		"`inventory`.`name` AS inventoryName ",
	)
	client.Joins("" +
		"JOIN `inventory_operation_status` ON " +
		"`inventory_operations`.`status` = `inventory_operation_status`.`id` " +
		"JOIN `user` u1 ON u1.`id` = `inventory_operations`.`src_executor` " +
		"JOIN `user` u2 ON u2.`id` = `inventory_operations`.`dst_executor`" +
		"JOIN `inventory_operations_detail` ON `inventory_operations_detail`.`operation_id` = `inventory_operations`.`id`" +
		"JOIN `inventory` ON `inventory_operations_detail`.`inventory_id` = `inventory`.`id`",
	)

	if len(filter.Ids) == 0 && len(filter.SrcExecutors) == 0 && len(filter.DstExecutors) == 0 &&
		len(filter.RequestTimes) == 0 && len(filter.StatusTimes) == 0 && len(filter.Statuses) == 0 {
		client.Find(&result)
	} else {
		if len(filter.Ids) != 0 {
			client.Where(filter.Ids)
		}

		if len(filter.SrcExecutors) != 0 {
			client.Where("src_executor IN (?)", filter.SrcExecutors)
		}

		if len(filter.DstExecutors) != 0 {
			client.Where("dst_executor IN (?)", filter.DstExecutors)
		}

		if len(filter.RequestTimes) != 0 {
			if len(filter.RequestTimes) != 2 {
				err = fmt.Errorf("request times should be 2 values")
			}

			a, _ := time.Parse("2006-01-02 15:04:05", filter.RequestTimes[0])
			b, _ := time.Parse("2006-01-02 15:04:05", filter.RequestTimes[1])

			client.Where("request_time BETWEEN ? AND ?", a.Format("2006-01-02 15:04:05.000000"), b.Format("2006-01-02 15:04:05.000000"))
		}

		if len(filter.StatusTimes) != 0 {
			if len(filter.StatusTimes) != 2 {
				err = fmt.Errorf("status times should be 2 values")
			}

			a, _ := time.Parse("2006-01-02 15:04:05", filter.StatusTimes[0])
			b, _ := time.Parse("2006-01-02 15:04:05", filter.StatusTimes[1])

			client.Where("request_time BETWEEN ? AND ?", a.Format("2006-01-02 15:04:05.000000"), b.Format("2006-01-02 15:04:05.000000"))
		}

		if len(filter.Statuses) != 0 {
			client.Where("inventory_operation_status.name IN (?)", filter.Statuses)
		}

		client.Find(&result)
	}

	return result, client.Error
}
