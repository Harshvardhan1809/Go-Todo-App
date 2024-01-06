package controllers

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
	"io"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/Harshvardhan1809/Go-Todo-App/utils"
	"github.com/Harshvardhan1809/Go-Todo-App/models"
)

var newTask models.Task

func GetUsers(w http.ResponseWriter, r *http.Request){

	//fmt.Println("In the get users controller")

	users := models.GetAllUsers()
	res, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTasks(w http.ResponseWriter, r *http.Request){

	//fmt.Println("In the get tasks controller")

	// Get the user ID
	vars := mux.Vars(r)
	userId := vars["user_id"]
	UserID, _ := strconv.ParseInt(userId, 0, 0)
	// if err != nil {
	// 	fmt.Println("Error while parsing")
	// }
	// Get all tasks
	newTasks := models.GetAllTasks(UserID)
	res, _ := json.Marshal(newTasks)

	// Send data over response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	// if error == nil {
	// 	fmt.Println("Printed the tasks")
	// } else if error != nil {
	// 	fmt.Println("Got an error ", error)
	// 	panic(error.Error())
	// }
}
 
func GetTaskByID(w http.ResponseWriter, r *http.Request){

	//fmt.Println("In the get task by id controller")

	// Get the date from the request, obtain object and convert to JSON
	vars := mux.Vars(r)
	taskId := vars["id"]
	ID, _ := strconv.ParseInt(taskId, 0, 0)
	// if err != nil {
	// 	fmt.Println("Error while parsing")
	// }
	taskDetails, _ := models.GetTaskByID(ID)
	res, _ := json.Marshal(taskDetails)

	// Send data over the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PostTaskNew(w http.ResponseWriter, r *http.Request){

	//fmt.Println("In the post new task controller")

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

	//fmt.Println("In the delete controller")

	vars := mux.Vars(r)
	taskId := vars["task_id"]
	ID, _ := strconv.ParseInt(taskId, 0, 0)
	// if err != nil {
	// 	fmt.Println("Error while parsing")
	// }

	t := models.DeleteTaskByID(ID)
	res, _ := json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	// if error != nil{
	// 	//fmt.Println("Getting err while deleting", error)
	// 	log.Fatal("Error deleting task")
	// 	panic(error.Error())
	// }
}

func GetTaskByDate(w http.ResponseWriter, r *http.Request){

	//fmt.Println("In the get task by date controller")

	vars := mux.Vars(r)
	date := vars["date"]

	d, _ := time.Parse("2006-01-02", date) // formatting to the ISO layout
	// if error != nil {
	// 	fmt.Println("Improper date format")
	// } else {
	// 	fmt.Println("Successfully parsed the date", d)
	// }

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
		fmt.Println("Error while parsing user id")
	}

	// YYYY-MM-DD 00:00:00 +0000 UTC -> YYYY-MM-DD 23:59:59 +0000 UTC
	d1, _ := time.Parse("2006-01-02", date) // formatting to the ISO layout
	// if error != nil {
	// 	fmt.Println("Improper date format")
	// } else {
	// 	fmt.Println("Successfully parsed the date", d1)
	// }
	d2 := d1.Add(23*time.Hour).Add(59*time.Minute).Add(59*time.Second)

	t := models.GetTaskByUserIDDate(ID, d1, d2)
	//fmt.Println("Print all the fetched tasks")
	//fmt.Println(t)
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	

}

func GetTaskPrevious(w http.ResponseWriter, r *http.Request){

	//fmt.Println("In the controller to get tasks from previous 5 days")

	vars := mux.Vars(r)
	userId := vars["user_id"]
	ID, _ := strconv.ParseInt(userId, 0, 0)
	// if err != nil {
	// 	fmt.Println("Error while parsing user id")
	// } 

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
	taskId, _ := strconv.ParseInt(task_id, 0, 0)
	// if err != nil {
	// 	fmt.Println("Error while parsing task id")
	// }

	// Get the task data and update it
	task, db := models.GetTaskByID(taskId)

	if newTask.CompletedAt != task.CompletedAt {
		task.CompletedAt = newTask.CompletedAt
		db.Model(&task).Updates(map[string]interface{}{"CompletedAt": newTask.CompletedAt, "Status": newTask.Status})
	} else if newTask.Starred != task.Starred {
		db.Model(&task).Updates(map[string]interface{}{"Starred": newTask.Starred})
	}

	res, _ := json.Marshal(task)
	//fmt.Println("All the damn way here")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}