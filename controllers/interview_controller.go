package controllers

import (
	"hoditgo/services"
	"hoditgo/services/models"
	"encoding/json"
	"net/http"
)

func CreateInterview(w http.ResponseWriter, r *http.Request) {
	interview := new(models.Interview)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&interview)
	services.CreateInterview(interview)
	w.Header().Set("Content-Type", "application/json")
}
