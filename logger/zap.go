package zap

import (
	"bytes"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
	"time"
)

type Zap struct {
	Directory      string // 目录
	LoggerFileName string //文件名
	DirectoryType  string //日志类型/等级
	Suffix         string //后缀

	Day         int64 // 最大保存天数（天）
	CuttingTime int64 // 按照时间切割（分钟）（默认1分钟）

	LoggerType bool // 输出日志类型 Console：true (默认);JSON：false
	ISConsole  bool // 是否输出到系统日志 默认（true）
}

func NewZapLogger(data *Zap) *Logger {

	split := strings.Split(data.DirectoryType, ",")
	dynamicLevel := zap.NewAtomicLevel()
	dynamicLevel.SetLevel(zap.InfoLevel)
	var path = make(map[string]string)

	for _, v := range split {
		path[v] = data.Directory + "/" + "%Y%m%d%H%M" + data.LoggerFileName + "-" + v + data.Suffix
	}

	var logger = new(Logger)
	if !data.LoggerType {

		var coreConfig = []CoreConfig{
			{
				Enc: zapcore.NewConsoleEncoder(coreEncoderConfig()),
				Ws:  zapcore.AddSync(os.Stdout),
				Lvl: dynamicLevel,
			},
			{
				Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["all"], data.Day, data.CuttingTime),
				Lvl: zap.NewAtomicLevelAt(zapcore.DebugLevel),
			},
			{
				Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["debug"], data.Day, data.CuttingTime),
				Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
					return lev == zap.DebugLevel
				}),
			},
			{
				Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["info"], data.Day, data.CuttingTime),
				Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
					return lev == zap.InfoLevel
				}),
			},
			{
				Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["warn"], data.Day, data.CuttingTime),
				Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
					return lev == zap.WarnLevel
				}),
			},
			{
				Enc: zapcore.NewJSONEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["error"], data.Day, data.CuttingTime),
				Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
					return lev >= zap.ErrorLevel
				}),
			},
		}
		//if data.ISConsole {
		//	coreConfig = append(coreConfig, CoreConfig{
		//		Enc: zapcore.NewConsoleEncoder(coreEncoderConfig()),
		//		Ws:  zapcore.AddSync(os.Stdout),
		//		Lvl: dynamicLevel,
		//	})
		//}
		logger = NewLogger(WithCores(coreConfig...),
			WithZapOptions(zap.AddCaller()), // 行号
			//WithExtraKeys([]ExtraKey{"requestId"}),
		)
	} else {
		var coreConfig = []CoreConfig{
			{
				Enc: zapcore.NewConsoleEncoder(coreEncoderConfig()),
				Ws:  zapcore.AddSync(os.Stdout),
				Lvl: dynamicLevel,
			},
			{
				Enc: zapcore.NewConsoleEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["all"], data.Day, data.CuttingTime),
				Lvl: zap.NewAtomicLevelAt(zapcore.DebugLevel),
			},
			{
				Enc: zapcore.NewConsoleEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["debug"], data.Day, data.CuttingTime),
				Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
					return lev == zap.DebugLevel
				}),
			},
			{
				Enc: zapcore.NewConsoleEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["info"], data.Day, data.CuttingTime),
				Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
					return lev == zap.InfoLevel
				}),
			},
			{
				Enc: zapcore.NewConsoleEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["warn"], data.Day, data.CuttingTime),
				Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
					return lev == zap.WarnLevel
				}),
			},
			{
				Enc: zapcore.NewConsoleEncoder(pkgEncoderConfig()),
				Ws:  getWriteSyncer(path["error"], data.Day, data.CuttingTime),
				Lvl: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
					return lev >= zap.ErrorLevel
				}),
			},
		}
		//if data.ISConsole {
		//	coreConfig = append(coreConfig, CoreConfig{
		//		Enc: zapcore.NewConsoleEncoder(coreEncoderConfig()),
		//		Ws:  zapcore.AddSync(os.Stdout),
		//		Lvl: dynamicLevel,
		//	})
		//}
		logger = NewLogger(WithCores(coreConfig...),
			WithZapOptions(zap.AddCaller()), // 行号
			//WithExtraKeys([]ExtraKey{"requestId"}),
		)
	}

	buf := new(bytes.Buffer)
	logger.SetOutput(buf)
	defer logger.Sync()

	return logger
}

// encoderConfig copy from zap
func pkgEncoderConfig() zapcore.EncoderConfig {
	cfg := coreEncoderConfig()
	cfg.EncodeTime = ISO8601TimeEncoder
	cfg.EncodeLevel = CapitalLevelEncoder
	cfg.EncodeDuration = StringDurationEncoder
	return cfg
}

// ISO8601TimeEncoder 自定义时间格式
func ISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, "2006-01-02 15:04:05", enc)
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {

	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}
	enc.AppendString("[" + t.Format(layout) + "]")
}

// CapitalLevelEncoder 自定义等级格式
func CapitalLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + l.CapitalString() + "]")
}

// StringDurationEncoder 自定义时间格式
func StringDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + d.String() + "]")
}

func ShortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func coreEncoderConfig() zapcore.EncoderConfig {
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
		EncodeCaller:   ShortCallerEncoder,
	}
}

/**
*	日志分割
*	file:文件名
*	day:最大保存天数
*	cuttingTime:按照时间进行切割
*	WithLinkName：软连接，用于日志切割
 */
func getWriteSyncer(file string, day, cuttingTime int64) zapcore.WriteSyncer {
	logf, err := rotatelogs.New(
		file,
		rotatelogs.WithLinkName(file),
		rotatelogs.WithMaxAge(time.Duration(day)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(cuttingTime)*time.Minute),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
	}
	return zapcore.AddSync(logf)
}
