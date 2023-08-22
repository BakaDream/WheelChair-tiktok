package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

var Logger *zap.Logger
var DSN = os.Getenv("MYSQL_DSN")
var DB *gorm.DB
