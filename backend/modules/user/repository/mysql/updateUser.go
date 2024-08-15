package mysql

import (
	"backend/internal/cmd/entity"
	"backend/modules/user/repository/mysql/models"
	"errors"
)

func (r *MysqlRepository) UpdateUser(user models.UserObject) error {
	client := r.client.Table(entity.User.Table)
	tx := client.Updates(&user)

	if tx.RowsAffected == 0 {
		return errors.New("NU")
	}

	return tx.Error
}
