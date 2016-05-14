package services

import (
	"hoditgo/core/repositories"
	"hoditgo/services/models"
	"net/http"
)


func CreateInterviewer(interviewer *models.Interviewer) (int, []byte) {
    interviewBackEnd := repositories.InitInterviewerRepo()
	interviewBackEnd.CreateInterviewer(interviewer)
	return http.StatusOK, []byte("")
}

func GetInterviewerById(id string) (models.Interviewer){
    interviewerBackEnd := repositories.InitInterviewerRepo()
    interviewer := interviewerBackEnd.GetInterviewerById(id)
    
    return interviewer
}

func GetInterviewerByName(name string) ([]models.Interviewer){
    interviewerBackEnd := repositories.InitInterviewerRepo()
    interviewer := interviewerBackEnd.GetInterviewerByName(name)
    
    return interviewer
}


func UpdateSummary(id, summary string)(int, []byte){
     interviewerBackEnd := repositories.InitInterviewerRepo()
     interviewerBackEnd.UpdateInterviewerSummary(summary, id)
     return http.StatusOK, []byte("")
}

func UpdateRanking(id string, ranking int)(int,[]byte){
    interviewerBackEnd := repositories.InitInterviewerRepo()
    interviewerBackEnd.UpdateInterviewerRanking(ranking,id)
    return http.StatusOK, []byte("")
}
