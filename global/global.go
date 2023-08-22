package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.Logger

var DB *gorm.DB
