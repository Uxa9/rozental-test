package models

type UserObjects []UserObject

type UsersObject struct {
	Ids   []int    `json:"ids"`
	Names []string `json:"names"`
}

type UserObject struct {
	Id   int    `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`
}
