package mysql

import (
	"backend/internal/cmd/entity"
	"backend/modules/status/repository/mysql/model"
)

func (r *MysqlRepository) CreateNewStatus(name string) error {
	client := r.client.Table(entity.Status.Table)
	status := model.StatusObject{Name: name}

	return client.Create(&status).Error
}
