package model

import "backend/modules/status/repository/mysql/model"

type GetStatusInput struct {
	Ids   []int    `json:"ids"`
	Names []string `json:"names"`
}

type GetStatusOutput struct {
	List  model.StatusObjects `json:"list"`
	Count int                 `json:"count"`
}
