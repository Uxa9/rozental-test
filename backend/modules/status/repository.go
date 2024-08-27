package status

import (
	"backend/modules/status/repository/mysql/model"
)

type MysqlRepoInterface interface {
	CreateNewStatus(name string) error
	GetStatus(filter model.StatusFilter) (model.StatusObjects, error)
}
