package models

type OperationStatus struct {
	StatusId   int    `json:"id"   gorm:"column:statusId"`
	StatusName string `json:"name" gorm:"column:nameStatus"`
}
