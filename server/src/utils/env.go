package utils

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvVariables(){

	err := godotenv.Load("/Users/harshvardhankedare/Desktop/Coding/golang/todoapp/server/src/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}