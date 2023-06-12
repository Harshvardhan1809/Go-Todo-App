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
	"github.com/gorilla/mux"
	"github.com/Harshvardhan1809/Go-Todo-App/controllers" 
)

var TodoAppRoutes = func(router *mux.Router){
	router.HandleFunc("/task/{date}", controllers.GetTaskByDate).Methods("GET")
}