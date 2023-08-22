package model

import (
	"database/sql"
	"fmt"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DSN = os.Getenv("MYSQL_DSN")
var TABLES = []interface{}{
	&Video{}, &Comment{}, &User{}, &UserVideoLike{}, &UserFollow{},
}
var DB *gorm.DB

func databaseCreate() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	config, err := mysqlDriver.ParseDSN(DSN) //解析
	if err != nil {
		fmt.Println("Failed to parse DSN:", err)
		return
	}
	// 提取数据库名称
	dbName := config.DBName
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MySQL database created successfully")
} //创建新的数据库

func databseInit() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	for _, table := range TABLES {
		db.AutoMigrate(table)
	}
	log.Println("MySQL database init successfully")
} //数据库初始化

func DatabaseConn() {
	databaseCreate()
	databseInit()
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
