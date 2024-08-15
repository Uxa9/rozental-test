package mysql

import "gorm.io/gorm"

type MysqlRepository struct {
	client *gorm.DB
}

func NewMysqlRepository(client *gorm.DB) *MysqlRepository {
	return &MysqlRepository{
		client: client,
	}
}
