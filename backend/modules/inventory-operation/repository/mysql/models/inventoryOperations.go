package models

import "time"

type OperationFilter struct {
	Ids          []int    `json:"id"`
	SrcExecutors []int    `json:"src_executor"`
	DstExecutors []int    `json:"dst_executor"`
	RequestTimes []string `json:"request_time"`
	StatusTimes  []string `json:"status_time"`
	Statuses     []string `json:"status"`
}

type Operation struct {
	Id          int       `gorm:"primary_key;column:id"`
	SrcExecutor int       `gorm:"column:src_executor"`
	DstExecutor int       `gorm:"column:dst_executor"`
	RequestTime time.Time `gorm:"column:request_time"`
	StatusTimes time.Time `gorm:"column:status_times"`
	Status      string    `gorm:"column:status"`
}
