package models

type InventoryFilters struct {
	Id         []int    `json:"id"`
	Name       []string `json:"name"`
	ExecutorId []int    `json:"executor_id"`
	Status     []string `json:"status"`
}

type Inventory struct {
	Id         int    `gorm:"primary_key;column:id"`
	Name       string `gorm:"column:name"`
	ExecutorId int    `gorm:"column:id_executor"`
}
