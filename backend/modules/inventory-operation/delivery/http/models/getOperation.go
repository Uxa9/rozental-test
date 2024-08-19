package models

import "time"

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

type OperationInventory struct {
	InventoryId   int    `json:"inventory_id"  gorm:"column:idInventory"`
	InventoryName string `json:"inventory"  gorm:"column:inventoryName"`
}

type InventoryOperation struct {
	DstExecutor OperationDstExecutor `json:"dst_executor" gorm:"foreignKey:DstExecutor;references:idDstExecutor;embedded"`
	RequestTime time.Time            `json:"request_time"`
	StatusTime  time.Time            `json:"status_time"`
	Status      OperationStatus      `json:"status" gorm:"foreignKey:Status;references:IdStatus;embedded"`
}

type InventoryOperationsDetails = []InventoryOperationDetails

type InventoryOperationDetails struct {
	Id          int `json:"id"`
	OperationId int `json:"operation_id"`
	InventoryId int `json:"inventory_id"`
}
