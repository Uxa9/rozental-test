package status

import "backend/modules/status/delivery/http/model"

type UseCaseInterface interface {
	CreateStatus(name string) error
	GetStatus(filter model.GetStatusInput) (model.GetStatusOutput, error)
}
