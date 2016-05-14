package controllers

import (
	"hoditgo/services"
	"hoditgo/services/models"
	"encoding/json"
	"net/http"
)

func CreateInterviewer(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	interview := new(models.Interviewer)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&interview)
	services.CreateInterviewer(interview)
	w.Header().Set("Content-Type", "application/json")
}

func UpdateInterviewerRanking(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	interviewer := new(models.Interviewer)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&interviewer)
	services.UpdateRanking(interviewer.Id, interviewer.Ranking)
	w.Header().Set("Content-Type", "application/json")
}


func UpdateInterviewerSummary(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	interviewer := new(models.Interviewer)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&interviewer)
	services.UpdateSummary(interviewer.Id, interviewer.Summary)
	w.Header().Set("Content-Type", "application/json")
}

func GetInterviewerById(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	var id string
	id = r.URL.Query().Get("InterviewerId")
	interviewer := services.GetInterviewerById(id)
	encoder := json.NewEncoder(w)
	encoder.Encode(interviewer)
}

func GetInterviewerByName(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	name := r.URL.Query().Get("InterviewerName")
	interviewer := services.GetInterviewerByName(name)
	encoder := json.NewEncoder(w)
	encoder.Encode(interviewer)
}

