package controllers

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
	"reflect"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/Harshvardhan1809/Go-Todo-App/utils"
	"github.com/Harshvardhan1809/Go-Todo-App/models"
)

var newTask models.Task

func GetUsers(w http.ResponseWriter, r *http.Request){

	fmt.Println("In the get users controller")

	users := models.GetAllUsers()
	res, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTasks(w http.ResponseWriter, r *http.Request){

	fmt.Println("In the get tasks controller")

	// Get the user ID
	vars := mux.Vars(r)
	userId := vars["user_id"]
	UserID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	// Get all tasks
	newTasks := models.GetAllTasks(UserID)
	res, _ := json.Marshal(newTasks)

	// Send data over response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, error := w.Write(res)

	if error == nil {
		fmt.Println("Printed the tasks")
	} else if error != nil {
		fmt.Println("Got an error ", error)
		panic(error.Error())
	}
}
 
func GetTaskByID(w http.ResponseWriter, r *http.Request){

	fmt.Println("In the get task by id controller")

	// Get the date from the request, obtain object and convert to JSON
	vars := mux.Vars(r)
	taskId := vars["id"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	taskDetails, _ := models.GetTaskByID(ID)
	res, _ := json.Marshal(taskDetails)

	// Send data over the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PostTaskNew(w http.ResponseWriter, r *http.Request){

	fmt.Println("In the post new task controller")

	r.ParseForm()
	for key, value := range r.Form {
		fmt.Println("%s = %s \n", key, value)
	}

	// task := &models.Task{} 
	// task.CreatedAt = time.Now()
	// utils.ParseBody(r, task)
	// t := task.CreateTask()
	// res, _ := json.Marshal(t)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

// HTTP ERROR 405 - can be solved if a proper request is sent
func DeleteTaskByID(w http.ResponseWriter, r *http.Request){

	fmt.Println("In the delete controller")

	vars := mux.Vars(r)
	taskId := vars["id"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	t := models.DeleteTaskByID(ID)
	res, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ , error := w.Write(res)

	if error == nil {
		fmt.Println("No error")
	}
	if error != nil{
		fmt.Println("Getting err while deleting", error)
		panic(error.Error())
	}
}

func GetTaskByDate(w http.ResponseWriter, r *http.Request){

	fmt.Println("In the get task by date controller")

	vars := mux.Vars(r)
	userId := vars["user_id"]
	date := vars["date"]

	fmt.Println(userId, date, reflect.TypeOf(date))

	ID, err := strconv.ParseInt(userId, 0, 0)
	if err == nil {
		fmt.Println("Error while parsing user id")
	}

	d, error := time.Parse("2006-01-02", date) // formatting to the ISO layout
	if error != nil {
		fmt.Println("Improper date format")
	} else {
		fmt.Println("Successfully parsed the date", d)
	}

	t := models.GetTaskByDate(ID, d)
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	

}

func GetTaskPrevious(w http.ResponseWriter, r *http.Request){

	fmt.Println("In the controller to get tasks from previous 5 days")

	vars := mux.Vars(r)
	userId := vars["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing user id")
	} 

	t := models.GetTaskPrevious(ID)
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// util.ParseBody parses the body of the request for us
func UpdateTaskByID(w http.ResponseWriter, r *http.Request){

	fmt.Println("In the controller to update tasks by id")

	// first store the contents of the request body in newTask
	var newTask = &models.Task{}
	utils.ParseBody(r, newTask)

	// Parse the task Id
	vars := mux.Vars(r)
	task_id := vars["id"]
	taskId, err := strconv.ParseInt(task_id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing task id")
	}

	// Get the task data and update it
	task, db := models.GetTaskByID(taskId)
	if newTask.Title != "" {
		task.Title = newTask.Title
	}
	if newTask.CompletedAt != task.CompletedAt {
		task.CompletedAt = newTask.CompletedAt
	}
	task.UpdatedAt = newTask.UpdatedAt
	task.Description = newTask.Description
	task.Status = newTask.Status
	task.Starred = newTask.Starred

	db.Save(&task)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}