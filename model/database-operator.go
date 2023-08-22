package model

import (
	"WheelChair-tiktok/global"
	"database/sql"
	"errors"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var TABLES = []interface{}{
	&Video{}, &Comment{}, &User{}, &UserVideoLike{}, &UserFollow{},
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

func LogUp(name string, password string) bool {
	var user User
	result := global.DB.Where("UserName = ?", name).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 没有找到记录
			hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal("Hash error\n")
			}
			user = User{Password: string(hash), UserName: name}
			global.DB.Create(&user)
			return true
		} else {
			// 查询过程中发生了其他错误
			log.Fatal("Unknown error")
		}
	} else {
		log.Println("The user already exists")
	}
	return false
}

func LogIn(name string, password string) bool {
	var user User
	result := global.DB.Where("UserName = ? AND Password = ?", name).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 没有找到记录
			log.Println("The username or password is incorrect")
			return false
		} else {
			// 查询过程中发生了其他错误
			log.Fatal("Unknown error\n")
		}
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Hash error\n")
		}
		if string(hash) == user.Password {
			log.Println("Login successful")
		}
	}
	return true
}
