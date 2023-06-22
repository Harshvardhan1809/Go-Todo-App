package main 

import (
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Harshvardhan1809/Go-Todo-App/routes"
	"github.com/Harshvardhan1809/Go-Todo-App/utils"
)


func init(){
	// function that runs before main	
	utils.LoadEnvVariables()	
	// use ENV variables using os.Getenv()
}

func main(){

	r := mux.NewRouter();
	// r.HandleFunc("/home", homeHandler)
	// http.ListenAndServe(":80", r)		

	routes.TodoAppRoutes(r)
	http.Handle("/", r)
	port := os.Getenv("PORT")
	url := "localhost:" + port
	log.Fatal(http.ListenAndServe(url, r))

} 
