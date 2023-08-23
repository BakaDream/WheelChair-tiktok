package model

import (
	"WheelChair-tiktok/global"
	"database/sql"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var TABLES = []interface{}{
	&Video{}, &Comment{}, &User{}, &UserVideoLike{}, //&UserFollow,
}

func databaseCreate() {
	db, err := sql.Open("mysql", global.DSN)
	if err != nil {
		log.Fatal("Failed to connect database\n")
	}
	defer db.Close()
	config, err := mysqlDriver.ParseDSN(global.DSN) //解析
	if err != nil {
		log.Fatal("Failed to parse DSN:", err)
		return
	}
	// 提取数据库名称
	dbName := config.DBName
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal("Failed to create database\n")
	}
	log.Println("MySQL database create successfully")
} //创建新的数据库

func databseInit() {
	db, err := gorm.Open(mysql.Open(global.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init database\n")
	}
	for _, table := range TABLES {
		db.AutoMigrate(table)
	}
	log.Println("MySQL database init successfully")
} //数据库初始化

func DatabaseConn() {
	databaseCreate()
	databseInit()
	db, err := gorm.Open(mysql.Open(global.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database\n")
	}
	log.Println("MySQL database connect successfully")
	global.DB = db
}
