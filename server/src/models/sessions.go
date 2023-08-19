package models

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
	"github.com/Harshvardhan1809/Go-Todo-App/config"
)

var (
	store *mysqlstore.MySQLStore
	session_name = "user_session"
 	session *sessions.Session
)

func init(){

	config.ConnectSessionStore()
	store := config.GetSessionStore();

	// Include the user struct in the session
	gob.Register(&User{})


	session = sessions.NewSession(store, session_name)
	store.Options = &sessions.Options{
		Domain: "192.168.102.222",
		Path: "/",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   true,
	}


}