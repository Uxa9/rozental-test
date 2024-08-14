package mysql

import (
	"backend/internal/cmd/entity"
	"backend/modules/user/repository/mysql/models"
)

func (r *MysqlRepository) CreateUser(name string) error {
	client := r.client.Table(entity.User.UserTable)
	user := models.UserObject{Name: name}
	return client.Create(&user).Error
}
