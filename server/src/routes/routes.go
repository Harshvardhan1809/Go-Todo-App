package routes

// API : Routes and request
// GET task by date : /task/{date}
// GET incomplete tasks from 5 previous days : /task/prev
// GET task by id : /task/{id}
// PUT task : /task/{id}
// POST new task : /task/
// DELETE task : /task/{id}

// POST sign in form : /auth/signin
// GET login form : /auth/login

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/Harshvardhan1809/Go-Todo-App/controllers" 
	"github.com/Harshvardhan1809/Go-Todo-App/middleware" 
)

var TodoAppRoutes = func(router *mux.Router){
	router.HandleFunc("/task/delete/{id}/", controllers.DeleteTaskByID).Methods("DELETE")
	router.HandleFunc("/task/user/{user_id}/", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/task/date/{user_id}/{date}/", controllers.GetTaskByDate).Methods("GET")
	router.HandleFunc("/task/prev/{user_id}/", controllers.GetTaskPrevious).Methods("GET")
	router.HandleFunc("/task/{id}/", controllers.GetTaskByID).Methods("GET")
	router.HandleFunc("/task/update/{id}/", controllers.UpdateTaskByID).Methods("PUT")
	router.HandleFunc("/task/{user_id}/", controllers.PostTaskNew).Methods("POST") 

	router.HandleFunc("/users/", controllers.GetUsers).Methods("GET")

	router.HandleFunc("/auth/signin", controllers.SignUp).Methods("POST")
	router.HandleFunc("/auth/login", controllers.LogIn).Methods("GET")
	router.HandleFunc("/auth/validate", middleware.RequireAuth).Methods("GET")

	fmt.Println("Didn't match any route lol")
}
// controllers.Validate