package models

import "time"

type InventoryFilters struct {
	Id         []int    `json:"id"`
	Name       []string `json:"name"`
	ExecutorId []int    `json:"executor_id"`
	Status     []string `json:"status"`
}

type InventoryOperationsDetails = []InventoryOperationDetails

type InventoryOperationDetails struct {
	Id          int `json:"id"`
	OperationId int `json:"operation_id"`
	InventoryId int `json:"inventory_id"`
}

type InventoryOperation struct {
	DstExecutor OperationDstExecutor `json:"dst_executor" gorm:"foreignKey:DstExecutor;references:idDstExecutor;embedded"`
	RequestTime time.Time            `json:"request_time"`
	StatusTime  time.Time            `json:"status_time"`
	Status      OperationStatus      `json:"status" gorm:"foreignKey:Status;references:IdStatus;embedded"`
}

type InventoryOperationResults = []InventoryOperationResult

type InventoryOperationResult struct {
	Id          int                  `json:"id"`
	Name        string               `json:"name"`
	Executor    OperationSrcExecutor `json:"executor" gorm:"foreignKey:Executor;references:idSrcExecutor;embedded"`
	Transaction InventoryOperation   `json:"transaction" gorm:"foreignKey:Transaction;references:transactionId;embedded"`
}

type InventoryOperationResultsNested = []InventoryOperationResultNested

type InventoryOperationResultNested struct {
	Id           int                  `json:"id"`
	Name         string               `json:"name"`
	Executor     OperationSrcExecutor `json:"executor"`
	Transactions []InventoryOperation `json:"transactions"`
}
