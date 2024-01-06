package main

import (
	"github.com/Harshvardhan1809/Go-Todo-App/routes"
	"github.com/Harshvardhan1809/Go-Todo-App/utils"
	"github.com/MadAppGang/httplog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
)

func init() {
	// function that runs before main
	utils.LoadEnvVariables()
	// use ENV variables using os.Getenv()
}

// To use Air for live reloading : alias air="$(go env GOPATH)"/bin/air

func main() {

	r := mux.NewRouter()
	r.Use(httplog.LoggerWithFormatter(httplog.DefaultLogFormatter))

	routes.TodoAppRoutes(r)
	http.Handle("/", r)
	port := os.Getenv("PORT")
	url := "localhost:" + port

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	credentialsOk := handlers.AllowCredentials()
	exposedHeaders := handlers.ExposedHeaders([]string{"Set-Cookie"})

	log.Fatal(http.ListenAndServe(url, handlers.CORS(headersOk, methodsOk, originsOk, credentialsOk, exposedHeaders)(r)))

}
