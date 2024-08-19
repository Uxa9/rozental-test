package models

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
