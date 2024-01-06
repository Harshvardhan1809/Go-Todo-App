package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"encoding/json"
	_ "io/ioutil"
	"github.com/golang-jwt/jwt/v4"
	"github.com/Harshvardhan1809/Go-Todo-App/config"
	"github.com/Harshvardhan1809/Go-Todo-App/models"
	"github.com/Harshvardhan1809/Go-Todo-App/utils"
)

type Claims struct {
	Username string `json:"usr"`
	Authorized bool `json:"aut"`
	jwt.RegisteredClaims
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func RequireAuth(w http.ResponseWriter, r *http.Request){

	// Get the cookies from request
	t, err := r.Cookie("token")
	if err != nil {		
		utils.FillErrorResponse(&w, http.StatusBadRequest, "Cookie not found, try again")
		return
	}
	tokenString := t.Value

	// Unhash it, check the expiry
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// This line checks if the token Method is HMAC, or else it can't proceed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		registerClaims := claims["RegisteredClaims"].(map[string]interface {})
		// have to do this since the type of
		// claims["RegisteredClaims"] is not known so we have to do type assertion
		expTime := registerClaims["exp"]
		if float64(time.Now().Unix()) > expTime.(float64){
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Expires: time.Now(),
			})
			w.WriteHeader(http.StatusOK)
		}

		// Find the user with the token
		var user models.User
		db := config.GetDB()
		db.Where("username=?", claims["Username"]).Find(&user)

		if user.ID == 0 {
			utils.FillErrorResponse(&w, http.StatusNotFound, "User not found, try again")
		}	
		
		// User found attach to request 
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	utils.FillErrorResponse(&w, http.StatusUnauthorized, "Unauthenticated, try again")
	return
}

func CheckSessionMiddleware (f http.HandlerFunc) http.HandlerFunc { 
	
	return func(w http.ResponseWriter, r *http.Request){

		// Get the cookies from request
		t, err := r.Cookie("token")
		if err != nil {
			log.Fatal("Cookie not found")
		}
		tokenString := t.Value
		claims := &Claims{};

		// Unhash it, check the expiry
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return  []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.FillErrorResponse(&w, http.StatusUnauthorized, "Error in authentication, try again")
				return
			}
			utils.FillErrorResponse(&w, http.StatusBadRequest, "Bad request, try again")
			return
		}
		if !token.Valid {
			utils.FillErrorResponse(&w, http.StatusUnauthorized, "Error in token, try again")
			return
		}

		if time.Until(claims.ExpiresAt.Time) > 0*time.Second {
			utils.FillErrorResponse(&w, http.StatusUnauthorized, "Token expiration")
			return
		}
		
		f(w,r) 
	}
}