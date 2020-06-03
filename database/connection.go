package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetMysqlConnection() *gorm.DB {
	usernameAndPassword := fmt.Sprint(os.Getenv("db_user")) + ":" + fmt.Sprint(os.Getenv("db_password"))
	hostName := "tcp(" + fmt.Sprint(os.Getenv("db_host")) + ":" + fmt.Sprint(os.Getenv("db_port")) + ")"

	log.Println("Connecting to DB Server " + fmt.Sprint(os.Getenv("db_host")) + ":" + fmt.Sprint(os.Getenv("db_port")) + "...")
	urlConnection := usernameAndPassword + "@" + hostName + "/" + fmt.Sprint(os.Getenv("db_database")) + "?charset=utf8&parseTime=true&loc=UTC"

	db, err := gorm.Open("mysql", urlConnection)
	if err != nil {
		panic(err)
	}

	log.Println("DB Server is connected!")
	return db
}
