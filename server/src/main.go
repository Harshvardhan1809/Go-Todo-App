package main 

import (
	_ "fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){

	r := mux.NewRouter();
	r.HandleFunc("/home", homeHandler)

	http.ListenAndServe(":80", r)		

} 
