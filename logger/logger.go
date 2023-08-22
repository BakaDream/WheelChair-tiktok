package logger

import (
	"WheelChair-tiktok/global"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// InitLogger 初始化Logger组件
func InitLogger() error {
	var level zapcore.Level

	// 解析日志级别
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		return err
	}

	var output zapcore.WriteSyncer

	// 根据日志输出类型创建WriteSyncer
	switch logOutput {
	case "stdout":
		output = zapcore.Lock(zapcore.AddSync(os.Stdout))
	case "stderr":
		output = zapcore.Lock(zapcore.AddSync(os.Stderr))
	default:
		// 如果需要将日志输出到文件，可以在这里添加相应的逻辑
		// output = zapcore.Lock(zapcore.AddSync(yourLogFile))
		return errors.New("unsupported log output type")
	}

	// 配置日志编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		output,
		level,
	)

	// 创建Logger
	global.Logger = zap.New(core)

	// 替换全局Logger
	zap.ReplaceGlobals(global.Logger)

	return nil
}
