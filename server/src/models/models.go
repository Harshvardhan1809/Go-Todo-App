package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Harshvardhan1809/Go-Todo-App/config"
	"time"
	"fmt"
	"reflect"
)

var db *gorm.DB

// db.Model(&user).Related(&emails)
//// SELECT * FROM emails WHERE user_id = 111; // 111 is user's primary key

type User struct{
	gorm.Model // by default includes ID, CreatedAt, DeletedAt and UpdatedAt fields
	Username string `json:"username" gorm:"unique"` 
	Password string `json:"-" gorm:"unique"`
	Tasks []Task `json:"tasks"` // Foreign key
}

type Task struct{ // for the values which allow null, we use omitempty
	gorm.Model
	UserID uint `json:"userId"` // foreign key
	CompletedAt time.Time `json:"completedAt,omitempty" `
	Title string `json:"name"`
	Description string `json:"description,omitempty"`
	Status string `json:"status,omitempty"`
	Starred bool `json:"starred,omitempty"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() (*User, *gorm.DB){
	db.NewRecord(u)
	db.Create(&u)
	return u, db
}

func GetAllUsers() []User{
	var Users []User
	db.Find(&Users)
	return Users
} 

// func GetUser() (*User, *gorm.DB){
// 	var user User
// 	db.First(&user, "email = ?", )
// 	return user, db
// }

func (t *Task) CreateTask() (*Task, *gorm.DB){
	db.NewRecord(t)
	db.Create(&t)
	return t, db
}

func GetAllTasks(UserId int64) []Task{
	var Tasks []Task
	db.Where("USER_ID=?", UserId).Find(&Tasks) 
	// basically .Find() does : SELECT * FROM Tasks and .Where is the WHERE condition
	return Tasks
}

func GetTaskByID(Id int64) (*Task, *gorm.DB){
	var getTask Task
	db := db.Where("ID=?", Id).Find(&getTask)
	fmt.Println("Printing the task we got", getTask, getTask.CreatedAt, reflect.TypeOf(getTask.CreatedAt))
	return &getTask, db
}

func DeleteTaskByID(Id int64) Task{
	var task Task
	var task_pre Task
	db.Where("ID=?", Id).Find(&task_pre)
	db.Where("ID=?", Id).Unscoped().Delete(task) // 
	return task
}

// Problem -> getting task even when a date is specified
func GetTaskByDate(UserId int64, d time.Time) []Task{	

	var Tasks []Task
	db.Where("created_at is null").Find(&Tasks)

	return Tasks

}

func GetTaskPrevious(UserId int64) []Task{

	var Tasks []Task
	db.Where("created_at is null").Find(&Tasks)

	return Tasks
}