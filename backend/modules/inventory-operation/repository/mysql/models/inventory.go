package models

type InventoryFilters struct {
	Id         []int    `json:"id"`
	Name       []string `json:"name"`
	ExecutorId []int    `json:"executor_id"`
	Status     []string `json:"status"`
}
