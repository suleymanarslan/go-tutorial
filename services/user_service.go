package services

import (
	"hoditgo/core/repositories"
	"hoditgo/services/models"
	"net/http"
)

func CreateUser(requestUser *models.User) (int, []byte) {
	userBackEnd := repositories.InitUserRepository()
	userBackEnd.CreateUser(requestUser)
	return http.StatusOK, []byte("")
}

func GetUser(requestUser *models.User) (models.User) {
	userBackEnd := repositories.InitUserRepository()
	user := userBackEnd.GetUser(requestUser)
	return user
}

