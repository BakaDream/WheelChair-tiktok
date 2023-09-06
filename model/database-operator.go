package model

import (
	l "WheelChair-tiktok/logger"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"
)

var DB *gorm.DB

func Init() {
	//生成dsn
	dsn, err := generateDSN()
	if err != nil {
		l.Logger.Fatal(err.Error())
		return
	}
	// 创建数据库如果 没有创建
	err = createDataBase()
	if err != nil {
		l.Logger.Fatal(err.Error())
		return
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		l.Logger.Fatal(err.Error())
		return
	}

	err = db.AutoMigrate(&User{}, &Video{}, &Favorite{}, &Follow{}, &Comment{})
	if err != nil {
		l.Logger.Fatal(err)
	}
	err = addDefaultUser(db)
	if err != nil {
		l.Logger.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		l.Logger.Fatal("Failed to get sqlDB")
		return
	}
	// 设置最大空闲连接数
	MaxIdleCount, _ := strconv.Atoi(os.Getenv("MAX_IDLE_COUNT"))
	MaxOpenCount, _ := strconv.Atoi(os.Getenv("MAX_OPEN"))
	MaxLifeTime, _ := strconv.Atoi(os.Getenv("MAX_LIFE_TIME"))
	MaxIdleTime, _ := strconv.Atoi(os.Getenv("MAX_IDLE_TIME"))
	sqlDB.SetMaxIdleConns(MaxIdleCount)
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(MaxOpenCount)
	// 设置连接的最大存活时间
	sqlDB.SetConnMaxLifetime(time.Duration(MaxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Duration(MaxIdleTime))

	l.Logger.Info("MySQL database connect successfully")
	DB = db
}

// 生成dsn
func generateDSN() (string, error) {
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	userPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 检查必要的环境变量是否已设置
	if mysqlHost == "" || mysqlPort == "" || dbName == "" || dbUser == "" || userPassword == "" {
		return "", errors.New("mysql params is invalid")
	}

	// 拼接 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", dbUser, userPassword, mysqlHost, mysqlPort, dbName)
	return dsn, nil
}

func createDataBase() error {
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	userPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUser, userPassword, mysqlHost, mysqlPort)
	dbc, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 创建数据库并设置字符集为 UTF-8 MB4
	_, err = dbc.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " CHARACTER SET utf8mb4")
	if err != nil {
		return err
	}
	dbc.Close()
	return nil
}
func addDefaultUser(db *gorm.DB) error {
	u := &User{
		Model:           gorm.Model{ID: 1},
		UserName:        "default",
		IP:              "",
		Password:        "asdbkasbdlasbd",
		FollowCount:     0,
		FollowerCount:   0,
		Signature:       "",
		Avatar:          "",
		BackgroundImage: "",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}
	// 判断用户是否注册
	if !errors.Is(db.Where("ID = ?", u.ID).First(&u).Error, gorm.ErrRecordNotFound) {
		return nil
	}

	// 否则注册
	err := db.Create(&u).Error
	return err
}
