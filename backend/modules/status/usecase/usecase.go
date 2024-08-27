package usecase

import (
	"backend/modules/status"
	"backend/modules/status/repository/mysql"
	"gorm.io/gorm"
)

type StatusUC struct {
	mysqlRepo status.MysqlRepoInterface
}

func NewStatusUseCase(db *gorm.DB) *StatusUC {
	return &StatusUC{
		mysqlRepo: mysql.NewMysqlRepository(db),
	}
}
