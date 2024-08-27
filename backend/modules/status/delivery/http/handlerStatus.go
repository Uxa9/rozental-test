package http

import (
	"backend/modules/status/delivery/http/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func ConvertQueryToArr(queryName string, context *gin.Context) []string {
	queryArr := context.QueryArray(queryName)

	if len(queryArr) != 0 {
		return strings.Split(strings.ReplaceAll(context.QueryArray(queryName)[0], " ", ""), ",")
	}

	return make([]string, 0)
}

func ConvertStringArrToInt(arr []string) []int {
	var idArr []int

	for _, id := range arr {
		intId, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		idArr = append(idArr, intId)
	}

	return idArr
}

func (handler *Handler) CreateStatus(context *gin.Context) {

	var status model.CreateStatusInput

	if err := context.ShouldBindJSON(&status); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.statusUC.CreateStatus(status.Name); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": status.Name})
	return
}

func (handler *Handler) GetStatus(context *gin.Context) {

	var statusArr model.GetStatusOutput
	var filter model.GetStatusInput
	idQuery := ConvertStringArrToInt(ConvertQueryToArr("Ids", context))
	nameQuery := ConvertQueryToArr("Names", context)

	filter.Ids = idQuery
	filter.Names = nameQuery

	statusArr, err := handler.statusUC.GetStatus(filter)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"status": statusArr})

	return
}
