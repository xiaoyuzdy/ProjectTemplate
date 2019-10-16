package log

import (
	"github.com/spf13/viper"
	"go-web/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime/debug"
)

var Sugar *zap.SugaredLogger

func InitLogs() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(string(debug.Stack()))
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	utils.MkDir("./runtime")
	config := zap.Config{
		Level:            atom,                           // 日志级别
		Development:      true,                           // 开发模式，堆栈跟踪
		Encoding:         "json",                         // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                  // 编码器配置
		OutputPaths:      []string{"./runtime/http.log"}, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	var isDebug = viper.GetBool("system.Debug")

	if isDebug {
		config.OutputPaths = []string{"stdout", "stderr"}
	}

	// 构建日志
	logger, _ := config.Build()
	Sugar = logger.Sugar()
}
