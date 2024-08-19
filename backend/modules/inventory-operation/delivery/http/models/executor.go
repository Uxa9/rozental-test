package models

type OperationSrcExecutor struct {
	ExecutorId   int    `json:"executor_id"  gorm:"column:idSrcExecutor"`
	ExecutorName string `json:"executor"`
}

type OperationDstExecutor struct {
	ExecutorId   int    `json:"executor_id"  gorm:"column:idDstExecutor"`
	ExecutorName string `json:"executor"`
}
