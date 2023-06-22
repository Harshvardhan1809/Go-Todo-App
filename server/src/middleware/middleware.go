package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	_ "io/ioutil"
	"github.com/golang-jwt/jwt/v4"
	"github.com/Harshvardhan1809/Go-Todo-App/config"
	"github.com/Harshvardhan1809/Go-Todo-App/models"
)

func RequireAuth(w http.ResponseWriter, r *http.Request){

	fmt.Println("In auth middleware")

	// Get the cookies from request
	t, err := r.Cookie("Authorization")
	if err != nil {
		log.Fatal("Cookie not found")
	}
	tokenString := t.Value

	// Unhash it, check the expiry
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			// Got the token
			// Check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64){
				log.Fatal("Token expired")
			}

			// Find the user with the token
			var user models.User
			db := config.GetDB()
			db.First(&user, claims["sub"])

			if user.ID == 0 {
				log.Fatal("User not found")
			}	

			// User found attach to request 
			w.Header().Set("UserId", string(user.ID))

	} else {
		fmt.Println(err)
	}

}