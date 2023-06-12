package controllers

import (
	"net/http"
	"github.com/Harshvardhan1809/Go-Todo-App/models"
)

var newTask models.Task
 
func GetTaskByDate(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Write(r)
}