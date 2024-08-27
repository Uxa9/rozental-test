package model

type StatusObjects []StatusObject

type StatusObject struct {
	Id   int    `gorm:"primary_key;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type StatusFilter struct {
	IdArr   []int
	NameArr []string
}
