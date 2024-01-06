package controllers

import (
	"net/http"
	"time"
	"strconv"
	"io"
	"encoding/json"
	"github.com/gorilla/mux"
 	"github.com/Harshvardhan1809/Go-Todo-App/utils"
	"github.com/Harshvardhan1809/Go-Todo-App/models"
)

var newTask models.Task

func GetUsers(w http.ResponseWriter, r *http.Request){

	users := models.GetAllUsers()
	res, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTasks(w http.ResponseWriter, r *http.Request){

	// Get the user ID
	vars := mux.Vars(r)
	userId := vars["user_id"]
	UserID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusBadRequest, "URL param should be an integer")
		return
	}

	// Get all tasks
	newTasks := models.GetAllTasks(UserID)
	res, _ := json.Marshal(newTasks)

	// Send data over response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
 
func GetTaskByID(w http.ResponseWriter, r *http.Request){

	// Get the date from the request, obtain object and convert to JSON
	vars := mux.Vars(r)
	taskId := vars["id"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusBadRequest, "URL param should be an integer")
		return
	}
	taskDetails, _ := models.GetTaskByID(ID)
	res, _ := json.Marshal(taskDetails)

	// Send data over the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PostTaskNew(w http.ResponseWriter, r *http.Request){

	newTask := &models.Task{}
	// CONVERSION
	// io.ReadCloser -> byte arrray -> string -> struct
	bufOfRequestBody, _ := io.ReadAll(r.Body)
	stringRequestBody := string(bufOfRequestBody)
	json.Unmarshal([]byte(stringRequestBody), &newTask)	

	t, _ := newTask.CreateTask()
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// HTTP ERROR 405 - can be solved if a proper request is sent
func DeleteTaskByID(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	taskId := vars["task_id"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusBadRequest, "URL param should be an integer")
		return
	}

	t := models.DeleteTaskByID(ID)
	res, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTaskByDate(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	date := vars["date"]

	d, err := time.Parse("2006-01-02", date) // formatting to the ISO layout
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusBadRequest, "URL param should be a proper date")
		return
	}

	t := models.GetTaskByDate(d)
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	

}

func GetTaskByUserIDDate(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	userId := vars["user_id"]
	date := vars["date"]

	//fmt.Println(userId, date, reflect.TypeOf(date))

	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusBadRequest, "URL param should be an integer")
		return
	}

	// YYYY-MM-DD 00:00:00 +0000 UTC -> YYYY-MM-DD 23:59:59 +0000 UTC
	// formatting to the ISO layout
	d1, _ := time.Parse("2006-01-02", date) 
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusBadRequest, "URL param should be a proper date")
		return
	}
	d2 := d1.Add(23*time.Hour).Add(59*time.Minute).Add(59*time.Second)

	t := models.GetTaskByUserIDDate(ID, d1, d2)
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	

}

func GetTaskPrevious(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	userId := vars["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusBadRequest, "URL param should be an integer")
		return
	}
	
	t := models.GetTaskPrevious(ID)
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// util.ParseBody parses the body of the request for us
func UpdateTaskByID(w http.ResponseWriter, r *http.Request){

	var newTask = &models.Task{}

	// CONVERSION
	// io.ReadCloser -> byte arrray -> string -> struct
	bufOfRequestBody, _ := io.ReadAll(r.Body)
	stringRequestBody := string(bufOfRequestBody)
	json.Unmarshal([]byte(stringRequestBody), &newTask)	

	// Parse the task Id
	vars := mux.Vars(r)
	task_id := vars["id"]
	taskId, err := strconv.ParseInt(task_id, 0, 0)
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusBadRequest, "URL param should be an integer")
		return
	}

	// Get the task data and update it
	task, db := models.GetTaskByID(taskId)

	if newTask.CompletedAt != task.CompletedAt {
		task.CompletedAt = newTask.CompletedAt
		db.Model(&task).Updates(map[string]interface{}{"CompletedAt": newTask.CompletedAt, "Status": newTask.Status})
	} else if newTask.Starred != task.Starred {
		db.Model(&task).Updates(map[string]interface{}{"Starred": newTask.Starred})
	}

	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}