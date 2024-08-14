package usecase

import (
	"backend/modules/user"
	"backend/modules/user/repository/mysql"
	"backend/modules/user/repository/mysql/models"
	"gorm.io/gorm"
)

type UserModuleUC struct {
	mysqlRepo user.MysqlRepoInterface
}

func NewUserModuleUseCase(mysqlRepo *gorm.DB) *UserModuleUC {
	return &UserModuleUC{
		mysqlRepo: mysql.NewMysqlRepository(mysqlRepo),
	}
}

func (uc *UserModuleUC) CreateUser(name string) error {
	return uc.mysqlRepo.CreateUser(name)
}

func (uc *UserModuleUC) GetUser(filter models.UsersObject) (models.UserObjects, error) {
	return uc.mysqlRepo.GetUser(filter)
}

func (uc *UserModuleUC) UpdateUser(user models.UserObject) error {
	return uc.mysqlRepo.UpdateUser(user)
}

func (uc *UserModuleUC) DeleteUser(id []int) error {
	return uc.mysqlRepo.DeleteUser(id)
}
