package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Harshvardhan1809/Go-Todo-App/config"
	"time"
)

var db *gorm.DB

// Define the schema for the model
// User
// Title 
// Description
// Date of creation
// Date of completion
// Status 
// Starred (boolean)

// db.Model(&user).Related(&emails)
//// SELECT * FROM emails WHERE user_id = 111; // 111 is user's primary key

type User struct{
	gorm.Model 
	Tasks []Task
}

type Task struct{
	gorm.Model
	UserID uint `json: "userId"` // foreign key
	CreatedAt time.Time `json: "createdAt"`
	CompletedAt time.Time `json: "completedAt"`
	Title string `json: "name"`
	Description string `json: "description"`
	Status string `json: "status"`
	Starred bool `json: "starred"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&User{})
}