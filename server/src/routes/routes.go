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

	fmt.Println("In the beginning of routes")

	// GET
	router.HandleFunc("/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/auth/check_auth", middleware.RequireAuth).Methods("GET")
	router.HandleFunc("/task/user/{user_id}/", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/task/date/{date}/", controllers.GetTaskByDate).Methods("GET")
	router.HandleFunc("/task/date/{user_id}/{date}/", controllers.GetTaskByUserIDDate).Methods("GET")
	router.HandleFunc("/task/prev/{user_id}/", controllers.GetTaskPrevious).Methods("GET")
	router.HandleFunc("/task/{id}/", controllers.GetTaskByID).Methods("GET")

	// POST
	router.HandleFunc("/auth/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	// router.HandleFunc("/auth/validate", middleware.RequireAuth).Methods("GET")
	router.HandleFunc("auth/session", controllers.CheckSession).Methods("GET")
	router.HandleFunc("/auth/logout", controllers.Logout).Methods("POST")
	router.HandleFunc("/task/{user_id}/", controllers.PostTaskNew).Methods("POST") 
	// router.HandleFunc("/auth/session", controllers.CheckSession).Methods("GET")

	// PUT
	router.HandleFunc("/task/update/{task_id}/", controllers.UpdateTaskByID).Methods("PUT")

	// DELETE
	router.HandleFunc("/task/delete/{task_id}/", controllers.DeleteTaskByID).Methods("DELETE")

	fmt.Println("Didn't match any route lol")
}
// controllers.Validate