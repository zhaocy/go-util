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

func loggerCore(fileName ,level string) []zapcore.Core{
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
		MaxSize:    64,  // 每个日志文件保存的最大尺寸 单位：M
		//MaxBackups: 30,                       // 日志文件最多保存多少个备份
		//MaxAge:     90,                        // 文件最多保存多少天
		LocalTime:  true,
		Compress:   false,  // 是否压缩
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
	return allCore
}

func InitServiceLogger(fileName string,level string, serviceName string) log.Logger{
	allCore:= loggerCore(fileName,level)
	// 设置初始化字段
	filed := zap.Fields(zap.String("service_name", serviceName))
	log := zap.New(zapcore.NewTee(allCore...), zap.AddCaller(), zap.AddCallerSkip(1),filed)
	gLogger = &zapLogger{
		logger: log,
		suger:  log.Sugar(),
	}
	return gLogger
}

func InitLogger(fileName string, level string) log.Logger {
	allCore:= loggerCore(fileName, level)
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

