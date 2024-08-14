package http

import (
	"backend/modules/user"
	"backend/modules/user/repository/mysql/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	userModuleUC user.UseCaseInterface
}

func NewHandler(userModuleUC user.UseCaseInterface) *Handler {
	return &Handler{
		userModuleUC: userModuleUC,
	}
}

func (handler *Handler) CreateUser(context *gin.Context) {

	type username struct {
		Username string `json:"username"`
	}

	var user username

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "username is empty"})
	}

	err := handler.userModuleUC.CreateUser(user.Username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"user": user.Username})

	return
}

func (handler *Handler) GetUser(context *gin.Context) {
	result := models.UserObjects{}
	idQuery := context.QueryArray("id")
	nameQuery := context.QueryArray("username")

	if len(idQuery) != 0 {
		idQuery = strings.Split(strings.ReplaceAll(context.QueryArray("id")[0], " ", ""), ",")
	}

	if len(nameQuery) != 0 {
		nameQuery = strings.Split(strings.ReplaceAll(context.QueryArray("username")[0], " ", ""), ",")
	}

	idsArr := []int{}

	for _, i := range idQuery {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		idsArr = append(idsArr, j)
	}

	filter := models.UsersObject{Ids: idsArr, Names: nameQuery}

	result, err := handler.userModuleUC.GetUser(filter)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, result)

	return
}

func (handler *Handler) UpdateUser(context *gin.Context) {

	var user models.UserObject

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Id <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id not set or bad value"})
	}

	if user.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "username is empty"})
	}

	err := handler.userModuleUC.UpdateUser(user)

	if err != nil {
		if err.Error() == "NU" {
			context.JSON(http.StatusOK, gin.H{"warning": "Nothing was changed, either you enter incorrect id or name was the same"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	context.JSON(http.StatusAccepted, gin.H{"user": user.Name})

	return
}

func (handler *Handler) DeleteUser(context *gin.Context) {
	idQuery := context.QueryArray("id")

	if len(idQuery) != 0 {
		idQuery = strings.Split(strings.ReplaceAll(context.QueryArray("id")[0], " ", ""), ",")
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id not set or bad value"})
	}

	idsArr := []int{}

	for _, i := range idQuery {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		idsArr = append(idsArr, j)
	}

	err := handler.userModuleUC.DeleteUser(idsArr)

	if err != nil {
		if err.Error() == "NF" {
			context.JSON(http.StatusOK, gin.H{"warning": "Nothing was changed, id not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	context.JSON(http.StatusOK, gin.H{"user(s) deleted id(s)": idsArr})

	return
}
