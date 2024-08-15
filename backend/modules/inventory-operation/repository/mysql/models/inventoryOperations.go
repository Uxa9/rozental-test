package models

import (
	"time"
)

type OperationFilter struct {
	Ids          []int    `json:"id"`
	SrcExecutors []int    `json:"src_executor"`
	DstExecutors []int    `json:"dst_executor"`
	RequestTimes []string `json:"request_time"`
	StatusTimes  []string `json:"status_time"`
	Statuses     []string `json:"status"`
}

type Operations []Operation

type Operation struct {
	Id          int       `gorm:"primary_key;column:id"`
	SrcExecutor int       `gorm:"column:src_executor"`
	DstExecutor int       `gorm:"column:dst_executor"`
	RequestTime time.Time `gorm:"column:request_time;default:CURRENT_TIMESTAMP()"`
	StatusTime  time.Time `gorm:"column:status_time;default:CURRENT_TIMESTAMP()"`
	Status      int       `gorm:"column:status"`
}

type OperationsResult []OperationResult

type OperationResult struct {
	Id          int                  `json:"id"`
	SrcExecutor OperationSrcExecutor `json:"src_executor" gorm:"foreignKey:SrcExecutor;references:idSrcExecutor;embedded"`
	DstExecutor OperationDstExecutor `json:"dst_executor" gorm:"foreignKey:DstExecutor;references:idDstExecutor;embedded"`
	Status      OperationStatus      `json:"status" gorm:"foreignKey:Status;references:IdStatus;embedded"`
	Inventory   OperationInventory   `json:"inventory" gorm:"foreignKey:Inventory;references:idInventory;embedded"`
	RequestTime time.Time            `json:"request_time"`
	StatusTime  time.Time            `json:"status_time"`
}

type OperationStatus struct {
	StatusId   int    `json:"id"   gorm:"column:idStatus"`
	StatusName string `json:"name" gorm:"column:nameStatus"`
}

type OperationSrcExecutor struct {
	ExecutorId   int    `json:"executor_id"  gorm:"column:idSrcExecutor"`
	ExecutorName string `json:"executor"  gorm:"column:nameSrcExecutor"`
}

type OperationDstExecutor struct {
	ExecutorId   int    `json:"executor_id"  gorm:"column:idDstExecutor"`
	ExecutorName string `json:"executor"  gorm:"column:nameDstExecutor"`
}

type OperationInventory struct {
	InventoryId   int    `json:"inventory_id"  gorm:"column:idInventory"`
	InventoryName string `json:"inventory"  gorm:"column:inventoryName"`
}
