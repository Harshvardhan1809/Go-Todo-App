package models

import (
	"github.com/Harshvardhan1809/Go-Todo-App/config"
)

func init(){

	config.ConnectSessionStore()
	// Include the user struct in the session
	// gob.Register(&User{})

}