package mysql

import (
	"backend/internal/cmd/entity"
	"backend/modules/user/repository/mysql/models"
)

func (r *MysqlRepository) GetUser(filters models.UsersObject) (result models.UserObjects, err error) {
	client := r.client.Table(entity.User.Table)

	if len(filters.Ids) == 0 && len(filters.Names) == 0 {
		client.Find(&result)
	} else {
		client.Where(filters.Ids).Or("name IN ?", filters.Names).Find(&result)
	}

	return result, client.Error
}
