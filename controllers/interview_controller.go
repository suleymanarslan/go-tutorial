package controllers

import (
	"hoditgo/services"
	"hoditgo/services/models"
	"encoding/json"
	"net/http"
)

func CreateInterview(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	interview := new(models.Interview)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&interview)
	services.CreateInterview(interview)
	w.Header().Set("Content-Type", "application/json")
}

func UpdateInterview(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	interview := new(models.Interview)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&interview)
	services.UpdateInterview(interview)
	w.Header().Set("Content-Type", "application/json")
}

func GetInterviewByName(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	var name string
	name = r.URL.Query().Get("name")
	interviewer := services.GetInterviewByName(name)
	encoder := json.NewEncoder(w)
	encoder.Encode(interviewer)
}