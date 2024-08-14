package user

import "backend/modules/user/repository/mysql/models"

type MysqlRepoInterface interface {
	CreateUser(name string) error
	GetUser(filter models.UsersObject) (models.UserObjects, error)
	UpdateUser(user models.UserObject) error
	DeleteUser(id []int) error
}
