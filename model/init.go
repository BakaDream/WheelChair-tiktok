package model

import (
	"gorm.io/gorm"
	"log"
	"os"
)
import "gorm.io/driver/mysql"

var DB *gorm.DB

func ConnDatabase() {

	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

}
