package main

import (
	_ "fmt"
	"net/http"
	_ "github.com/gorilla/mux"
	"html/template"
)

	// CREATE A STRUCT TO GIVE STRUCTURE TO THE WEBPAGE

type Todo struct {
	Title string
	Status bool
}

type Todolist struct {

	Header string
	Todos []Todo

} 

func homeHandler(w http.ResponseWriter, r *http.Request){

	// NOW DISPLAY THE HOME PAGE
	tmpl := template.Must(template.ParseFiles("../../client/templates/home.html", "../../client/templates/components.tmp"))

	// ADD DATA TO THE TEMPLATE
	data := Todolist {
		Header : "Harsh's todolist",
		Todos : []Todo{
			{ Title: "Complete LINE interview", Status: false },
			{ Title: "Prepare for 実験", Status: false },
			{ Title: "Go to gym ", Status: false},
		},
	}

	// EXECUTE THE TEMPLATE
	// execute applies the parsed template to the specificed data object
	// write the output to reponse using ResponseWriter
	tmpl.Execute(w, data)

}
