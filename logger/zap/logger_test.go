package zap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"strings"
	"testing"
)

// testEncoderConfig encoder config for testing, copy from zap
func testEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// humanEncoderConfig copy from zap
func humanEncoderConfig() zapcore.EncoderConfig {
	cfg := testEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeDuration = zapcore.StringDurationEncoder
	return cfg
}

//func getWriteSyncer(file string) zapcore.WriteSyncer {
//	logf, err := rotatelogs.New(
//		file,
//		rotatelogs.WithLinkName(file),
//		rotatelogs.WithMaxAge(180*24*time.Hour),
//		rotatelogs.WithRotationTime(time.Second*60),
//		//rotatelogs.WithRotationSize(500),
//	)
//	if err != nil {
//		log.Printf("failed to create rotatelogs: %s", err)
//	}
//	return zapcore.AddSync(logf)
//}

// TestLogger test logger work with hertz
func TestLogger(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewLogger(WithZapOptions(zap.WithFatalHook(zapcore.WriteThenPanic)))
	defer logger.Sync()

	hlog.SetLogger(logger)
	hlog.SetOutput(buf)
	hlog.SetLevel(hlog.LevelDebug)

	type logMap map[string]string

	logTestSlice := []logMap{
		{
			"logMessage":       "this is a trace log",
			"formatLogMessage": "this is a trace log: %s",
			"logLevel":         "Trace",
			"zapLogLevel":      "debug",
		},
		{
			"logMessage":       "this is a debug log",
			"formatLogMessage": "this is a debug log: %s",
			"logLevel":         "Debug",
			"zapLogLevel":      "debug",
		},
		{
			"logMessage":       "this is a info log",
			"formatLogMessage": "this is a info log: %s",
			"logLevel":         "Info",
			"zapLogLevel":      "info",
		},
		{
			"logMessage":       "this is a notice log",
			"formatLogMessage": "this is a notice log: %s",
			"logLevel":         "Notice",
			"zapLogLevel":      "warn",
		},
		{
			"logMessage":       "this is a warn log",
			"formatLogMessage": "this is a warn log: %s",
			"logLevel":         "Warn",
			"zapLogLevel":      "warn",
		},
		{
			"logMessage":       "this is a error log",
			"formatLogMessage": "this is a error log: %s",
			"logLevel":         "Error",
			"zapLogLevel":      "error",
		},
		{
			"logMessage":       "this is a fatal log",
			"formatLogMessage": "this is a fatal log: %s",
			"logLevel":         "Fatal",
			"zapLogLevel":      "fatal",
		},
	}

	testHertzLogger := reflect.ValueOf(logger)

	for _, v := range logTestSlice {
		t.Run(v["logLevel"], func(t *testing.T) {
			if v["logLevel"] == "Fatal" {
				defer func() {
					assert.Equal(t, "this is a fatal log", recover())
				}()
			}
			logFunc := testHertzLogger.MethodByName(v["logLevel"])
			logFunc.Call([]reflect.Value{
				reflect.ValueOf(v["logMessage"]),
			})
			assert.Contains(t, buf.String(), v["logMessage"])
			assert.Contains(t, buf.String(), v["zapLogLevel"])

			buf.Reset()

			logfFunc := testHertzLogger.MethodByName(fmt.Sprintf("%sf", v["logLevel"]))
			logfFunc.Call([]reflect.Value{
				reflect.ValueOf(v["formatLogMessage"]),
				reflect.ValueOf(v["logLevel"]),
			})
			assert.Contains(t, buf.String(), fmt.Sprintf(v["formatLogMessage"], v["logLevel"]))
			assert.Contains(t, buf.String(), v["zapLogLevel"])

			buf.Reset()

			ctx := context.Background()
			ctxLogfFunc := testHertzLogger.MethodByName(fmt.Sprintf("Ctx%sf", v["logLevel"]))
			ctxLogfFunc.Call([]reflect.Value{
				reflect.ValueOf(ctx),
				reflect.ValueOf(v["formatLogMessage"]),
				reflect.ValueOf(v["logLevel"]),
			})
			assert.Contains(t, buf.String(), fmt.Sprintf(v["formatLogMessage"], v["logLevel"]))
			assert.Contains(t, buf.String(), v["zapLogLevel"])

			buf.Reset()
		})
	}
}

// TestLogLevel test SetLevel
func TestLogLevel(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewLogger()
	defer logger.Sync()

	// output to buffer
	logger.SetOutput(buf)

	logger.Debug("this is a debug log")
	assert.False(t, strings.Contains(buf.String(), "this is a debug log"))

	logger.SetLevel(hlog.LevelDebug)

	logger.Debugf("this is a debug log %s", "msg")
	assert.True(t, strings.Contains(buf.String(), "this is a debug log"))

	logger.SetLevel(hlog.LevelError)
	logger.Infof("this is a debug log %s", "msg")
	assert.False(t, strings.Contains(buf.String(), "this is a info log"))

	logger.Warnf("this is a warn log %s", "msg")
	assert.False(t, strings.Contains(buf.String(), "this is a warn log"))

	logger.Error("this is a error log %s", "msg")
	assert.True(t, strings.Contains(buf.String(), "this is a error log"))
}

func TestWithCoreEnc(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewLogger(WithCoreEnc(zapcore.NewConsoleEncoder(humanEncoderConfig())))
	defer logger.Sync()

	// output to buffer
	logger.SetOutput(buf)

	logger.Infof("this is a info log %s", "msg")
	assert.True(t, strings.Contains(buf.String(), "this is a info log"))
}

func TestWithCoreWs(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewLogger(WithCoreWs(zapcore.AddSync(buf)))
	defer logger.Sync()

	logger.Infof("this is a info log %s", "msg")
	assert.True(t, strings.Contains(buf.String(), "this is a info log"))
}

func TestWithCoreLevel(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewLogger(WithCoreLevel(zap.NewAtomicLevelAt(zapcore.WarnLevel)))
	defer logger.Sync()

	// output to buffer
	logger.SetOutput(buf)

	logger.Infof("this is a info log %s", "msg")
	assert.False(t, strings.Contains(buf.String(), "this is a info log"))

	logger.Warnf("this is a warn log %s", "msg")
	assert.True(t, strings.Contains(buf.String(), "this is a warn log"))
}

// TestCoreOption test zapcore config option
func TestCoreOption(t *testing.T) {
	//loggera, _ := zap.NewProduction()
	//defer loggera.Sync()
	//loggera.Info("failed to fetch url",
	//	// 强类型字段
	//	zap.String("url", "http://example.com"),
	//	zap.Int("attempt", 3),
	//	zap.Duration("duration", time.Second),
	//)
	//loggera.With(
	//	// 强类型字段
	//	zap.String("url", "http://development.com"),
	//	zap.Int("attempt", 4),
	//	zap.Duration("duration", time.Second*5),
	//).Info("[With] failed to fetch url")
	//buf := new(bytes.Buffer)
	logger := NewZapLogger(&Zap{
		Directory:      "./logs",               // 目录
		LoggerFileName: "/system",              //文件名
		DirectoryType:  "all,info,debug,error", //日志类型/等级
		Suffix:         ".log",                 //后缀

		Day:         180, // 最大保存天数（天）
		CuttingTime: 1,   // 按照时间切割（分钟）

		LoggerType: true, // 输出日志类型 Console/JSON
		ISConsole:  true, // 是否输出到系统日志
	})
	ctx := context.WithValue(context.Background(), ExtraKey("requestId"), "123")
	//logger.Logger().Sugar().Info(123123)
	hlog.SetLogger(logger)
	hlog.CtxInfof(ctx, "123")

	//// test log level
	//assert.False(t, strings.Contains(buf.String(), "this is a debug log"))
	//
	//logger.Error("this is a warn log")
	//// test log level
	//assert.True(t, strings.Contains(buf.String(), "this is a warn log"))
	//// test console encoder result
	//assert.True(t, strings.Contains(buf.String(), "\tERROR\t"))
	//
	//logger.SetLevel(hlog.LevelDebug)
	//logger.Debug("this is a debug log")
	//assert.True(t, strings.Contains(buf.String(), "this is a debug log"))
	//
	////time.Sleep(time.Second * 66)
	//
	//logger.CtxInfof(ctx, "%s log", "extra")
	//// test log level
	//assert.False(t, strings.Contains(buf.String(), "this is a debug log"))
	//
	//logger.Error("this is a warn log")
	//// test log level
	//assert.True(t, strings.Contains(buf.String(), "this is a warn log"))
	//// test console encoder result
	//assert.True(t, strings.Contains(buf.String(), "\tERROR\t"))
	//
	//logger.SetLevel(hlog.LevelDebug)
	//logger.Debug("this is a debug log")
	//assert.True(t, strings.Contains(buf.String(), "this is a debug log"))
}

// TestCoreOption test zapcore config option
func TestZapOption(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := NewLogger(
		WithZapOptions(zap.AddCaller()),
	)
	defer logger.Sync()

	logger.SetOutput(buf)

	logger.Debug("this is a debug log")
	assert.False(t, strings.Contains(buf.String(), "this is a debug log"))

	logger.Error("this is a warn log")
	// test caller in log result
	assert.True(t, strings.Contains(buf.String(), "caller"))
}

// TestWithExtraKeys test WithExtraKeys option
func TestWithExtraKeys(t *testing.T) {
	buf := new(bytes.Buffer)

	log := NewLogger(WithExtraKeys([]ExtraKey{"requestId"}))
	log.SetOutput(buf)

	ctx := context.WithValue(context.Background(), ExtraKey("requestId"), "123")

	log.CtxInfof(ctx, "%s log", "extra")

	var logStructMap map[string]interface{}

	err := json.Unmarshal(buf.Bytes(), &logStructMap)

	assert.Nil(t, err)

	value, ok := logStructMap["requestId"]

	assert.True(t, ok)
	assert.Equal(t, value, "123")
}

func BenchmarkNormal(b *testing.B) {
	buf := new(bytes.Buffer)
	log := NewLogger()
	log.SetOutput(buf)
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		log.CtxInfof(ctx, "normal log")
	}
}

func BenchmarkWithExtraKeys(b *testing.B) {
	buf := new(bytes.Buffer)
	log := NewLogger(WithExtraKeys([]ExtraKey{"requestId"}))
	log.SetOutput(buf)
	ctx := context.WithValue(context.Background(), ExtraKey("requestId"), "123")
	for i := 0; i < b.N; i++ {
		log.CtxInfof(ctx, "normal log")
	}
}
