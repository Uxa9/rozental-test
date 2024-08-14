package mysql

import (
	"backend/internal/cmd/entity"
	"backend/modules/user/repository/mysql/models"
	"errors"
)

func (r *MysqlRepository) DeleteUser(id []int) error {
	client := r.client.Table(entity.User.UserTable)
	tx := client.Delete(models.UserObject{}, &id)

	if tx.RowsAffected == 0 {
		return errors.New("NF")
	}
	
	return tx.Error
}
