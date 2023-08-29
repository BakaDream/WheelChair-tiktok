package model

import (
	"WheelChair-tiktok/logger"
	"database/sql"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var DSN string

var TABLES = []interface{}{
	&Video{}, &Comment{}, &User{}, &Follow{}, &Favorite{},
}

func databaseCreate() {
	config, err := mysqlDriver.ParseDSN(DSN) //解析
	if err != nil {
		log.Fatal("Failed to parse DSN:", err)
		return
	}
	dbName := config.DBName
	newdbName := ""
	// 添加新的数据库名到配置中
	config.DBName = newdbName
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal("Failed to connect database\n")
	}
	defer db.Close()
	// 提取数据库名称
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal("Failed to create database\n")
	}
	log.Println("MySQL database create successfully")
} //创建新的数据库

func databseInit() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init database\n")
	}
	for _, table := range TABLES {
		db.AutoMigrate(table)
	}
	log.Println("MySQL database init successfully")
} //数据库初始化

func DatabaseConn() {
	DSN = os.Getenv("MYSQL_DSN")
	databaseCreate()
	databseInit()
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database\n")
	}
	logger.Logger.Info("MySQL database connect successfully")
	DB = db
}
