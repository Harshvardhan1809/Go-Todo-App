package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Harshvardhan1809/Go-Todo-App/routes"
)

func main(){

	r := mux.NewRouter();
	// r.HandleFunc("/home", homeHandler)
	// http.ListenAndServe(":80", r)		

	routes.TodoAppRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

} 
