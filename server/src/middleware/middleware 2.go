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
)

type Claims struct {
	Username string `json:"usr"`
	Authorized bool `json:"aut"`
	jwt.RegisteredClaims
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

// checks for jwt token in the response
func RequireAuth(w http.ResponseWriter, r *http.Request){

	//fmt.Println("In auth middleware")

	// Get the cookies from request
	t, _ := r.Cookie("token")
	//fmt.Println("Print the cookie ", t)
	// if err != nil {
	// 	fmt.Println("Cookie not found in the middleware")
	// }
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

		//fmt.Println("Claims checking ", claims);
		// Got the token
		// Check the exp
		registerClaims := claims["RegisteredClaims"].(map[string]interface {}) // have to do this since the type of
		// claims["RegisteredClaims"] is not known so we have to do type assertion
		expTime := registerClaims["exp"]
		if float64(time.Now().Unix()) > expTime.(float64){
			log.Fatal("Token expired")
		}

		// Find the user with the token
		var user models.User
		db := config.GetDB()
		db.Where("username=?", claims["Username"]).Find(&user)

		if user.ID == 0 {
			log.Fatal("User not found")
		}	
		
		// User found attach to request 
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}
	// } else {
	// 	fmt.Println(err)
	// }

	w.WriteHeader(http.StatusUnauthorized)
	//fmt.Println("User not found in the claims so we are sending error")
	return
}

func CheckSessionMiddleware (f http.HandlerFunc) http.HandlerFunc { 
	
	return func(w http.ResponseWriter, r *http.Request){

		//fmt.Println("In check session middleware")

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
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if time.Until(claims.ExpiresAt.Time) > 0*time.Second {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
		f(w,r) 
	}
}