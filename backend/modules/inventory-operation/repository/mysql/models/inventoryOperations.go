package models

type OperationFilter struct {
	Ids          []int    `json:"id"`
	SrcExecutors []int    `json:"src_executor"`
	DstExecutors []int    `json:"dst_executor"`
	RequestTimes []string `json:"request_time"`
	StatusTimes  []string `json:"status_time"`
	Statuses     []string `json:"status"`
}
