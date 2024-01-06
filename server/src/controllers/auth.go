package controllers

import (
	"encoding/json"
	"github.com/Harshvardhan1809/Go-Todo-App/config"
	"github.com/Harshvardhan1809/Go-Todo-App/models"
	"github.com/Harshvardhan1809/Go-Todo-App/utils"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"io"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Claims struct {
	Username   string `json:"usr"`
	Authorized bool   `json:"aut"`
	jwt.RegisteredClaims
}

func Signup(w http.ResponseWriter, r *http.Request) {

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
		utils.FillErrorResponse(&w, http.StatusInternalServerError, "Error during login, try again")
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

	// EXPIRE THE COOKIE, WRITE TO RESPONSE
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)

	// WRITE TO RESPONSE
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {

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
	db := config.GetDB()
	qError := db.Where("username=?", formBody.Username).Find(&user)
	fmt.Print("")
	if qError.Error != nil {
		utils.FillErrorResponse(&w, http.StatusNotFound, "User does not exist, create user first")
		return
	}

	// COMPARE INPUT PASSWORD WITH HASHED PASSWORD
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formBody.Password))
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusBadRequest, "Incorrect password, enter again")
		return
	}

	// CREATE JWT TOKEN - signing, putting in the request
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expirationTime := time.Now().Add(3 * time.Hour)
	claims["Authorized"] = true
	claims["Username"] = formBody.Username
	claims["RegisteredClaims"] = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	jwtKey := os.Getenv("SECRET")
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		utils.FillErrorResponse(&w, http.StatusInternalServerError, "Error during login, try again")
		return
	}

	// WRITE TO RESPONSE
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Secure:  true,
		// HttpOnly: true, enabling this would require using HTTPS all the time
		SameSite: http.SameSiteNoneMode,
	})
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	// EXPIRE THE COOKIE, WRITE TO RESPONSE
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
}
