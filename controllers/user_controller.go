package controllers

import (
	"hoditgo/services"
	"hoditgo/services/models"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
	services.CreateUser(requestUser)
	w.Header().Set("Content-Type", "application/json")
}
