package mysql

import (
	"backend/internal/cmd/entity"
	"backend/modules/status/repository/mysql/models"
)

func (r *MysqlRepository) CreateStatus(name string) error {
	client := r.client.Table(entity.Status.Table)
	user := models.StatusObject{Name: name}

	return client.Create(&user).Error
}
