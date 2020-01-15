package go_util

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"go.uber.org/zap"

	"github.com/go-log/log"
	"fmt"
	"strings"
	"os"
	"time"
)

type zapLogger struct {
	logger *zap.Logger
	suger *zap.SugaredLogger
}

func (s *zapLogger) Log(v ...interface{}) {
	s.suger.Warn(v...)
}

func (s *zapLogger) Logf(format string, v ...interface{}) {
	s.suger.Warnf(fmt.Sprintf(format, v...))
}

var gLogger *zapLogger

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05"))
}

func InitLogger(fileName string, level string) log.Logger {
	lv := zap.DebugLevel
	switch strings.ToLower(level) {
	case "info":
		lv = zap.InfoLevel
	case "warn":
		lv = zap.WarnLevel
	case "error":
		lv = zap.ErrorLevel
	case "panic":
		lv = zap.PanicLevel
	case "fatal":
		lv = zap.FatalLevel
	}

	var allCore []zapcore.Core
	syncWritter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    128,
		LocalTime:  true,
		Compress:   false,
	})
	enc := zap.NewProductionEncoderConfig()
	enc.EncodeTime = timeEncoder
	coreFile := zapcore.NewCore(zapcore.NewJSONEncoder(enc), syncWritter, zap.NewAtomicLevelAt(lv))
	allCore = append(allCore, coreFile)
	{
		consoleDebugging := zapcore.Lock(os.Stdout)
		conf := zap.NewDevelopmentEncoderConfig()
		conf.EncodeTime = timeEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(conf)

		coreConsole := zapcore.NewCore(consoleEncoder, consoleDebugging, zap.NewAtomicLevelAt(lv))
		allCore = append(allCore, coreConsole)
	}

	log := zap.New(zapcore.NewTee(allCore...), zap.AddCaller(), zap.AddCallerSkip(1))
	gLogger = &zapLogger{
		logger: log,
		suger:  log.Sugar(),
	}
	return gLogger
}

func Debug(args ...interface{}) {
	if gLogger == nil { return }
	gLogger.suger.Debug(args...)
}
func Debugf(format string ,args ...interface{}) {
	if gLogger == nil { return }
	gLogger.suger.Debugf(format, args...)
}
func Info(args ...interface{}) {
	if gLogger == nil { return }
	gLogger.suger.Info(args...)
}
func Infof(format string ,args ...interface{}) {
	if gLogger == nil { return }
	gLogger.suger.Infof(format, args...)
}
func Warn(args ...interface{}) {
	if gLogger == nil { return }
	gLogger.suger.Warn(args...)
}
func Warnf(format string ,args ...interface{}) {
	if gLogger == nil { return }
	gLogger.suger.Warnf(format, args...)
}
func Error(args ...interface{}) {
	if gLogger == nil { return }
	gLogger.suger.Error(args...)
}
func Errorf(format string ,args ...interface{}) {
	if gLogger == nil { return }
	gLogger.suger.Errorf(format, args...)
}
func Panic(args ...interface{}) {
	if gLogger == nil {
		panic(args)
		return
	}
	gLogger.suger.Panic(args...)
}
func Panicf(format string ,args ...interface{}) {
	if gLogger == nil {
		panic(args)
		return
	}
	gLogger.suger.Panicf(format, args...)
}
func Fatal(args ...interface{}) {
	if gLogger == nil {
		panic(args)
		return
	}
	gLogger.suger.Fatal(args...)
}
func Fatalf(format string ,args ...interface{}) {
	if gLogger == nil {
		panic(args)
		return
	}
	gLogger.suger.Panicf(format, args...)
}

