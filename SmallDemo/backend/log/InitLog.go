package log

import (
	"backend/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLog() {
	appConfig := config.GetAppConfig()

	lowestLevel := getLevel(appConfig.Logger.LowestLevel)
	logConfig := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "ts",
		CallerKey:    "file",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: func(duration time.Duration, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendInt64(int64((duration) / 1000000))
		},
	}

	var cores []zapcore.Core
	var levels = []string{"Debug", "Info", "Warn", "Error"}
	for _, l := range levels {
		currLevel := getLevel(l)
		if currLevel >= lowestLevel {
			enableLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
				return level == currLevel
			})
			logPath := getLogPath(currLevel)
			writer := getWriter(logPath)
			core := zapcore.NewCore(zapcore.NewConsoleEncoder(logConfig), zapcore.AddSync(writer), enableLevel)
			cores = append(cores, core)
		}
	}

	// log to Console
	cores = append(cores, zapcore.NewCore(
		zapcore.NewJSONEncoder(logConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		lowestLevel))

	core := zapcore.NewTee(cores...)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(lowestLevel))
}

func getLevel(lowestLevel string) zapcore.Level {
	switch lowestLevel {
	case "Debug":
		return zapcore.DebugLevel
	case "Info":
		return zapcore.InfoLevel
	case "Warn":
		return zapcore.WarnLevel
	case "Error":
		return zapcore.ErrorLevel
	}
	return zapcore.DebugLevel
}

func getLogPath(level zapcore.Level) string {
	switch level {
	case zapcore.DebugLevel:
		return config.GetAppConfig().Logger.Path.Debug
	case zapcore.InfoLevel:
		return config.GetAppConfig().Logger.Path.Info
	case zapcore.WarnLevel:
		return config.GetAppConfig().Logger.Path.Warn
	case zapcore.ErrorLevel:
		return config.GetAppConfig().Logger.Path.Error
	}
	return "./log/runtime/log.log"
}

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,
		MaxBackups: 2,
		MaxAge:     2,
		Compress:   false,
	}
}
