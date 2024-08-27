package mysql

import (
	"backend/internal/cmd/entity"
	"backend/modules/status/repository/mysql/model"
)

func (r *MysqlRepository) GetStatus(filter model.StatusFilter) (result model.StatusObjects, err error) {
	client := r.client.Table(entity.Status.Table)

	if len(filter.IdArr) == 0 && len(filter.NameArr) == 0 {
		client.Find(&result)
	} else {
		client.Where(filter.IdArr).Or("name IN ?", filter.NameArr).Find(&result)
	}

	return result, client.Error
}
