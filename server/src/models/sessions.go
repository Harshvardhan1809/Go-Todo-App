package models

import (
	_ "encoding/gob"
	_ "github.com/gorilla/sessions"
	_ "github.com/srinathgs/mysqlstore"
	_ "github.com/Harshvardhan1809/Go-Todo-App/config"
	_ "fmt"
)

// var (
// 	store *mysqlstore.MySQLStore
// 	session_name = "user_session"
//  	session *sessions.Session
// )

// func init(){

// 	fmt.Println("Print from sessions.go file")

// 	// config.ConnectSessionStore()
// 	// store := config.GetSessionStore();

// 	// // Include the user struct in the session
// 	// gob.Register(&User{})


// 	// session = sessions.NewSession(store, session_name)
// 	// store.Options = &sessions.Options{
// 	// 	Domain: "192.168.102.222",
// 	// 	Path: "/",
// 	// 	MaxAge:     60*60*60*2, // age is given in seconds
// 	// 	Secure:     false,
// 	// 	HttpOnly:   true,
// 	// }

// }