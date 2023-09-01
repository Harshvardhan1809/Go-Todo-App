package controllers

import (
	"fmt"
	"encoding/json"
	"time"
	"log"
	"io"
	"bytes"
	_ "os"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
	"github.com/Harshvardhan1809/Go-Todo-App/models"
	"github.com/Harshvardhan1809/Go-Todo-App/utils"
	"github.com/Harshvardhan1809/Go-Todo-App/config"
	"net/http"
)

var (
	store *mysqlstore.MySQLStore
	session *sessions.Session
)

func CheckSession(w http.ResponseWriter, r *http.Request){

	fmt.Println("Print from the controller", w);

	// utils.EnableCors(&w)

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

	utils.EnableCors(&w)
	fmt.Println("Print request body in signup ", r.Body)
	fmt.Println(" ")

	// Get the form data 
	var formBody struct {
		Username string
		Password string
	}
	utils.ParseBody(r, formBody)
	
	var newUser models.User
	var tasks []models.Task 
	newUser.Username = formBody.Username
	newUser.CreatedAt = time.Now()

	fmt.Println("Print the form body ", formBody)
	fmt.Println("Credentials ", newUser.Username, formBody.Password)

	// Hash the password 
	hash, err := bcrypt.GenerateFromPassword([]byte(formBody.Password), 10)
	if err != nil {
		log.Fatal("Error hashing the password")
	}
	newUser.Password = string(hash)
	newUser.Tasks = tasks

	fmt.Println("Credentials after hashing", newUser.Username, newUser.Password)

	// Create the user
	u, _ := newUser.CreateUser()
	res, _ := json.Marshal(u)

	fmt.Println("user ", u, res);

	// Respond
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request){

	fmt.Println("In the auth controller")

	// Get the email and password from the request body
	var formBody struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}
	// fmt.Println("Print the request in the controller", r.Body, formBody)
	// utils.ParseBody(r, formBody)

	bufOfRequestBody, _ := io.ReadAll(r.Body)
	fmt.Println("print the request body ", bufOfRequestBody);
	// [For Request Body] 消費されてしまったRequest Bodyを修復する
	r.Body = io.NopCloser(bytes.NewBuffer(bufOfRequestBody))

	fmt.Println("bufOfRequestBody : ", bytes.NewBuffer(bufOfRequestBody))
	fmt.Println("r.Body : ", r.Body)

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(formBody.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)	
		return
	}

	formBody.Password = string(hash)
	// [For Request Body] 消費されてしまったRequest Bodyを修復する
	r.Body = io.NopCloser(bytes.NewBuffer(bufOfRequestBody))
	fmt.Println("Print the password ", formBody.Password)

	// Search in DB
	var user models.User
	db := config.GetDB() //  , 
	qError := db.Where("username=?",formBody.Username).Where("password=?",formBody.Password).Find(&user)
	if qError.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// This is giving error
	// session, _ := store.Get(r, "session-name")
	// fmt.Println("Printing session", session)
	// session.Values["authenticated"] = true
	// session.Values["user"] = user
	// session.Save(r, w)

	res, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

