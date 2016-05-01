package services

import (
	"hoditgo/core/repositories"
	"hoditgo/services/models"
	"net/http"
)

func CreateInterview(interview *models.Interview) (int, []byte) {
	interviewBackEnd := repositories.InitInterviewRepository()
	interviewBackEnd.CreateInterview(interview)
	return http.StatusOK, []byte("")
}
