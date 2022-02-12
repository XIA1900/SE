package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLog(lowestLogLevel zapcore.Level) {
	config := zapcore.EncoderConfig{
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

	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.WarnLevel && level >= lowestLogLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel && level >= lowestLogLevel
	})

	infoPath := "./log/runtime/info.log"
	infoWriter := getWriter(infoPath)
	warnPath := "./log/runtime/warn.log"
	warnWriter := getWriter(warnPath)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(warnWriter), warnLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), lowestLogLevel),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
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
