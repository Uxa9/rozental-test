package http

import "backend/modules/status"

type Handler struct {
	statusUC status.UseCaseInterface
}

func NewHandler(statusUC status.UseCaseInterface) *Handler {
	return &Handler{statusUC: statusUC}
}
