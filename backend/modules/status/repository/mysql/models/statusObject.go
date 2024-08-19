package models

type StatusObjects []StatusObject

type StatusObjectFilter struct {
	Ids   []int    `json:"ids"`
	Names []string `json:"names"`
}

type StatusObject struct {
	Id   int    `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`
}
