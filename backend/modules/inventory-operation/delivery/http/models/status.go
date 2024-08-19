package models

type OperationStatus struct {
	StatusId   int    `json:"id"   gorm:"column:idStatus"`
	StatusName string `json:"name" gorm:"column:nameStatus"`
}
