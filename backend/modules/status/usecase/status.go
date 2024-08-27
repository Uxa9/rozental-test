package usecase

import (
	"backend/modules/status/delivery/http/model"
	sqlModel "backend/modules/status/repository/mysql/model"
)

func (uc *StatusUC) CreateStatus(name string) error {
	return uc.mysqlRepo.CreateNewStatus(name)
}

func (uc *StatusUC) GetStatus(filter model.GetStatusInput) (result model.GetStatusOutput, err error) {

	statusFilter := sqlModel.StatusFilter{
		IdArr:   filter.Ids,
		NameArr: filter.Names,
	}

	result.List, err = uc.mysqlRepo.GetStatus(statusFilter)
	result.Count = len(result.List)

	if err != nil {
		return model.GetStatusOutput{}, err
	}

	return result, err
}
