package main

import (
	"fmt"
	"reflect"
	"time"
)

// Proper datetime format compatible with MySQL 
// golangDateTime := time.Now().Format("2006-01-02 15:04:05") 

func m() {
	fmt.Println("Hello World")
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("Current date : ", t)
	fmt.Println("Type : ", reflect.TypeOf(t))

	dstr := "2023-06-17"
	d, err := time.Parse("2006-01-02", dstr)
	if err != nil {
		fmt.Println("Error !")
	} else {
		fmt.Println("")
		fmt.Println("Formatted date : ", d)
		fmt.Println("Type : ", reflect.TypeOf(d))
	}
}
