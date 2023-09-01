package controllers

import (
	"fmt"
	"encoding/json"
	"time"
	"log"
	"io"
	_ "os"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
	"github.com/Harshvardhan1809/Go-Todo-App/models"
	"github.com/Harshvardhan1809/Go-Todo-App/config"
	"net/http"
)

var (
	store *mysqlstore.MySQLStore
	session *sessions.Session
)

func CheckSession(w http.ResponseWriter, r *http.Request){

	fmt.Println("Print from the controller", w);

	store = config.GetSessionStore();
	session, err := store.Get(r, "session-name")
	if err != nil {
		fmt.Println("Error - could not get a session with the name")
	}

	// check if authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// get the user from the session
	user, _ := session.Values["data"].(*models.User)
	res, _ := json.Marshal(user)

	// return the user in the session
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Signup(w http.ResponseWriter, r *http.Request){

	// STRUCT FOR THE FORM BODY
	var formBody struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

 	// CONVERSION
	// io.ReadCloser -> byte arrray -> string -> struct
	bufOfRequestBody, _ := io.ReadAll(r.Body)
	stringRequestBody := string(bufOfRequestBody)
	json.Unmarshal([]byte(stringRequestBody), &formBody)

	// HASH THE PASSWORD AND ASSIGN TO NEW USER
	hash, err := bcrypt.GenerateFromPassword([]byte(formBody.Password), 10)
	if err != nil {
		log.Fatal("Error hashing the password")
	}

	// CREATE A NEW USER
	var newUser models.User
	var tasks []models.Task 
	newUser.Username = formBody.Username
	newUser.CreatedAt = time.Now()
	newUser.Tasks = tasks
	newUser.Password = string(hash)

	// CREATE THE USER
	u, _ := newUser.CreateUser()
	res, _ := json.Marshal(u)

	// WRITE TO RESPONSE
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request){

	// STRUCT FOR THE FORM BODY
	var formBody struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

	// CONVERSION
	// io.ReadCloser -> byte arrray -> string -> struct
	bufOfRequestBody, _ := io.ReadAll(r.Body)
	stringRequestBody := string(bufOfRequestBody)
	json.Unmarshal([]byte(stringRequestBody), &formBody)

	// SEARCH IN DB 
	var user models.User
	db := config.GetDB() //  , 
	qError := db.Where("username=?",formBody.Username).Find(&user)
	if qError.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// COMPARE INPUT PASSWORD WITH HASHED PASSWORD
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formBody.Password))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// This is giving error
	// session, _ := store.Get(r, "session-name")
	// fmt.Println("Printing session", session)
	// session.Values["authenticated"] = true
	// session.Values["user"] = user
	// session.Save(r, w)

	// WRITE TO RESPONSE
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
