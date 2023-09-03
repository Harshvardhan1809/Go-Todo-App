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
	"github.com/Harshvardhan1809/Go-Todo-App/models"
	"github.com/Harshvardhan1809/Go-Todo-App/config"
	_ "github.com/gorilla/sessions"
	"net/http"
)

// var (
//     // key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
//     key = []byte("super-secret-key")
//     store = sessions.NewCookieStore(key)
// )

func CheckSession(w http.ResponseWriter, r *http.Request){

	// store.GET creates a new session every time smh, we need cookie information from the mf request
	store := config.GetSessionStore();
	session, _ := store.Get(r, "k")
	fmt.Println("Print session in check session ", session.Values)

	// CHECK IF THE USER EXISTS IN SESSION
	if session.Values["username"] != "" {
		fmt.Println("Get user from DB ", session.Values["username"])
		var user models.User
		db := config.GetDB()
		qError := db.Where("username=?", session.Values["username"]).Find(&user)
		if qError.Error != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		res, _ := json.Marshal(user)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	// USER NOT FOUND, ERROR
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	return

	// check if authenticated
	// if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

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

	// CHECK SESSION
	store := config.GetSessionStore();
	session, getErr := store.Get(r, formBody.Username)
	if getErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("In the login controller")
	session.Values["username"] = formBody.Username
	session.Values["authenticated"] = true
	fmt.Println("Print the session values ", session.Values)
	fmt.Println("Print the session username ", session.Values["username"])
	fmt.Println("Print the session options ", session.Options.MaxAge)
	fmt.Println("Print the session name ", session.Name())
	session.Options.SameSite = http.SameSiteLaxMode;
	session.Options.HttpOnly = false;
	session.Options.Path = "/";
	session.Options.MaxAge = 60*60*3; //3hrs
	sessErr := session.Save(r, w) // saves this session in the store
	if sessErr != nil {
		fmt.Println("Can't store session in store")
		http.Error(w, sessErr.Error(), http.StatusInternalServerError)
		return
	}

	// ONLY THING REMAINING IS SENDING THE 

	// WRITE TO RESPONSE
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(user)
	fmt.Println("Print the response", string(res))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Logout (w http.ResponseWriter, r *http.Request) {

	// GET THE USERNAME FROM THE REQ BODY
	// io.ReadCloser -> []byte array -> string -> struct
	var logoutBody struct {
		Username string `json:"username,omitempty"`
	}
	bufOfRequestBody, _ := io.ReadAll(r.Body)
	stringRequestBody := string(bufOfRequestBody)
	json.Unmarshal([]byte(stringRequestBody), &logoutBody);

	// CHECK SESSION AND RETURN 200
	store := config.GetSessionStore();
	session, _ := store.Get(r, "session-name")
	if session.Values["username"] == logoutBody.Username {
		session.Options.MaxAge = -1;
		session.Save(r, w);
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

	// IF NO MATCH THEN RETURN 400 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	return
}