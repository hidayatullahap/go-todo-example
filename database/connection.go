package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetMysqlConnection() *gorm.DB {
	usernameAndPassword := fmt.Sprintf("%s:%s", os.Getenv("db_user"), os.Getenv("db_password"))
	hostName := fmt.Sprintf("tcp(%s:%s)", os.Getenv("db_host"), os.Getenv("db_port"))
	urlConnection := fmt.Sprintf("%s@%s/%s?charset=utf8&parseTime=true&loc=UTC", usernameAndPassword, hostName, os.Getenv("db_database"))

	log.Println(fmt.Sprintf("Connecting to DB Server %s:%s/%s...", os.Getenv("db_host"), os.Getenv("db_port"), os.Getenv("db_database")))
	db, err := gorm.Open("mysql", urlConnection)
	if err != nil {
		log.Println("DB Server failed to connect")
		panic(err)
	}

	log.Println("DB Server is connected!")
	return db
}
