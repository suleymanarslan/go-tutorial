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

func GetInterviews(offset int) ([]models.Interview){
    interviewBackEnd := repositories.InitInterviewRepository()
    interviews := interviewBackEnd.GetAllInterviews(offset)
    
    return interviews
}

func GetInterviewById(id string)(models.Interview){
     interviewBackEnd := repositories.InitInterviewRepository()
     return interviewBackEnd.GetInterviewById(id)
}

func DeactivateInterview(id string)(int, []byte){
     interviewBackEnd := repositories.InitInterviewRepository()
     interviewBackEnd.DeactivateInterview(id)
     return http.StatusOK, []byte("")
}

func UpdateInterview(interview *models.Interview)(int, []byte){
    interviewBackEnd := repositories.InitInterviewRepository()
    interviewBackEnd.UpdateInterview(interview)
    return http.StatusOK, []byte("")
}

func GetInterviewByName(name string)([]models.Interview){
    interviewBackEnd := repositories.InitInterviewRepository()
    return interviewBackEnd.GetInterviewByName(name)
}