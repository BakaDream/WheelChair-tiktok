package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var Logger *zap.SugaredLogger

// InitLogger 初始化Logger组件
func Init() {

	var logLevel zapcore.Level
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		{
			logLevel = zapcore.DebugLevel
		}
	case "INFO":
		{
			logLevel = zapcore.InfoLevel
		}
	default:
		log.Fatal("logLevel设置错误")
	}

	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logLevel)
	Logger = zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	//todo 彩色
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilepath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".log"

	// 日志切割
	maxSize, err := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	if err != nil {
		log.Fatal(err)
	}
	maxBackups, err := strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
	if err != nil {
		log.Fatal(err)
	}
	maxAge, err := strconv.Atoi(os.Getenv("LOG_MAX_AGE"))
	if err != nil {
		log.Fatal(err)
	}
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilepath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge, //days
		Compress:   true,   // disabled by default
	}
	return zapcore.AddSync(lumberjackSyncer)

}
