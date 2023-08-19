package config 

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/srinathgs/mysqlstore"
)

var (
	db *gorm.DB
	store *mysqlstore.MySQLStore
)

func ConnectDB(){
	d, err := gorm.Open("mysql", "root:MySQL@root@password@2002@/golang_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println("Error while opening database")
		// panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
} 

func ConnectSessionStore(){
	s, err := mysqlstore.NewMySQLStore("root:MySQL@root@password@2002@/golang_api?charset=utf8&parseTime=True&loc=Local", "sessions", "/", 3600, []byte(os.Getenv("SECRET")))
	if err != nil {
		fmt.Println("Error in connecting session store")
		panic(err)
	}
	store = s
}

func GetSessionStore() *mysqlstore.MySQLStore {
	return store
}