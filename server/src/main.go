package main 

import (
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/MadAppGang/httplog"
	"github.com/Harshvardhan1809/Go-Todo-App/routes"
	"github.com/Harshvardhan1809/Go-Todo-App/utils"
	"github.com/gorilla/handlers"
)

func init(){
	// function that runs before main	
	utils.LoadEnvVariables()	
	// use ENV variables using os.Getenv()
}

func main(){

	r := mux.NewRouter();
	r.Use(httplog.LoggerWithFormatter(httplog.DefaultLogFormatter))	
	
	routes.TodoAppRoutes(r)
	http.Handle("/", r)
	port := os.Getenv("PORT")
	url := "localhost:" + port 

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"});
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"});
	originsOk := handlers.AllowedOrigins([]string{"*"});

	log.Fatal(http.ListenAndServe(url, handlers.CORS(headersOk, methodsOk, originsOk)(r)))

} 
