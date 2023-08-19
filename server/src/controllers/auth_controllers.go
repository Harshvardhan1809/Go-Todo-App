package controllers

import (
	"fmt"
	"encoding/json"
	"time"
	"log"
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
	store = config.GetSessionStore();
	session, err := store.Get(r, "session-name")
	if err != nil {
		fmt.Println("Error - could not get a session with the name")
	}
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		fmt.Println("cant find authentication")
		return
	}
	// user, ok := session.Values["data"].(*sessions.User)

}

func SignUp(w http.ResponseWriter, r *http.Request){

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
	
	// Hash the password 
	hash, err := bcrypt.GenerateFromPassword([]byte(formBody.Password), 10)
	if err != nil {
		log.Fatal("Error hashing the password")
	}
	newUser.Password = string(hash)
	newUser.Tasks = tasks

	// Create the user
	u, _ := newUser.CreateUser()
	res, _ := json.Marshal(u)

	// Respond
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LogIn(w http.ResponseWriter, r *http.Request){

	// Get the email and password from the request body
	var formBody struct {
		Username string
		Password string
	}
	utils.ParseBody(r, formBody)

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(formBody.Password), 10)
	if err != nil {
		log.Fatal("Error hashing the password")
	}
	formBody.Password = string(hash)

	// Search in DB
	var user models.User
	db := config.GetDB()
	qError := db.Where("username=?",formBody.Username).Where("password=?",formBody.Password).Find(&user)
	if qError.Error != nil {
		log.Fatal("Error : User does not exist")
	}

	// Once user exists, we place it in the session
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Save(r, w)


	// Now that we have user, generate the token and hash it
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"sub" : user.ID,
	// 	"exp" : time.Now().Add(time.Hour * 24 * 365).Unix(),
	// })
	// tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	// if err != nil {
	// 	log.Fatal("Error signing the token")
	// }
	res, _ := json.Marshal(user)

	// // Set as cookie
	// cookie := http.Cookie{
	// 	Name: "Authorization", 
	// 	Value: tokenString, 
	// 	MaxAge: 3600 * 24 * 365,
	// 	Path:  "", 
	// 	Domain : "", 
	// 	Secure : false,
	// 	HttpOnly : true,
	// }
	// http.SetCookie(w, &cookie)

	// Login in using the token
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("token", tokenString)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

