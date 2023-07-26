package klogzap

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

var _ klog.FullLogger = (*Logger)(nil)

type Logger struct {
	l      *zap.Logger
	config *config
}

func NewLogger(opts ...Option) *Logger {
	config := new(config)
	// apply options
	for _, opt := range opts {
		opt.apply(config)
	}

	cores := make([]zapcore.Core, 0, len(config.coreConfigs))
	for _, coreConfig := range config.coreConfigs {
		cores = append(cores, zapcore.NewCore(coreConfig.Enc, coreConfig.Ws, coreConfig.Lvl))
	}
	logger := zap.New(
		zapcore.NewTee(cores[:]...),
		config.zapOpts...,
	)

	return &Logger{
		l:      logger,
		config: config,
	}
}

func (l *Logger) Log(level klog.Level, kvs ...interface{}) {
	sugar := l.l.Sugar().With()
	switch level {
	case klog.LevelTrace, klog.LevelDebug:
		sugar.Debug(kvs...)
	case klog.LevelInfo:
		sugar.Info(kvs...)
	case klog.LevelNotice, klog.LevelWarn:
		sugar.Warn(kvs...)
	case klog.LevelError:
		sugar.Error(kvs...)
	case klog.LevelFatal:
		sugar.Fatal(kvs...)
	default:
		sugar.Warn(kvs...)
	}
}

func (l *Logger) Logf(level klog.Level, format string, kvs ...interface{}) {
	logger := l.l.Sugar().With()
	switch level {
	case klog.LevelTrace, klog.LevelDebug:
		logger.Debugf(format, kvs...)
	case klog.LevelInfo:
		logger.Infof(format, kvs...)
	case klog.LevelNotice, klog.LevelWarn:
		logger.Warnf(format, kvs...)
	case klog.LevelError:
		logger.Errorf(format, kvs...)
	case klog.LevelFatal:
		logger.Fatalf(format, kvs...)
	default:
		logger.Warnf(format, kvs...)
	}
}

func (l *Logger) CtxLogf(level klog.Level, ctx context.Context, format string, kvs ...interface{}) {
	log := l.l.Sugar()
	if len(l.config.extraKeys) > 0 {
		for _, k := range l.config.extraKeys {
			log = log.With(string(k), ctx.Value(k))
		}
	}
	switch level {
	case klog.LevelDebug, klog.LevelTrace:
		log.Debugf(format, kvs...)
	case klog.LevelInfo:
		log.Infof(format, kvs...)
	case klog.LevelNotice, klog.LevelWarn:
		log.Warnf(format, kvs...)
	case klog.LevelError:
		log.Errorf(format, kvs...)
	case klog.LevelFatal:
		log.Fatalf(format, kvs...)
	default:
		log.Warnf(format, kvs...)
	}
}

func (l *Logger) Trace(v ...interface{}) {
	l.Log(klog.LevelTrace, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Log(klog.LevelDebug, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.Log(klog.LevelInfo, v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.Log(klog.LevelNotice, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.Log(klog.LevelWarn, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.Log(klog.LevelError, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Log(klog.LevelFatal, v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Logf(klog.LevelTrace, format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logf(klog.LevelDebug, format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logf(klog.LevelInfo, format, v...)
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	l.Logf(klog.LevelWarn, format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logf(klog.LevelWarn, format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logf(klog.LevelError, format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logf(klog.LevelFatal, format, v...)
}

func (l *Logger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelDebug, ctx, format, v...)
}

func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelDebug, ctx, format, v...)
}

func (l *Logger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelInfo, ctx, format, v...)
}

func (l *Logger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelWarn, ctx, format, v...)
}

func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelWarn, ctx, format, v...)
}

func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(klog.LevelFatal, ctx, format, v...)
}

func (l *Logger) SetLevel(level klog.Level) {
	var lvl zapcore.Level
	switch level {
	case klog.LevelTrace, klog.LevelDebug:
		lvl = zap.DebugLevel
	case klog.LevelInfo:
		lvl = zap.InfoLevel
	case klog.LevelWarn, klog.LevelNotice:
		lvl = zap.WarnLevel
	case klog.LevelError:
		lvl = zap.ErrorLevel
	case klog.LevelFatal:
		lvl = zap.FatalLevel
	default:
		lvl = zap.WarnLevel
	}

	l.config.coreConfigs[0].Lvl = lvl

	cores := make([]zapcore.Core, 0, len(l.config.coreConfigs))
	for _, coreConfig := range l.config.coreConfigs {
		cores = append(cores, zapcore.NewCore(coreConfig.Enc, coreConfig.Ws, coreConfig.Lvl))
	}

	logger := zap.New(
		zapcore.NewTee(cores[:]...),
		l.config.zapOpts...)

	l.l = logger
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.config.coreConfigs[0].Ws = zapcore.AddSync(writer)

	cores := make([]zapcore.Core, 0, len(l.config.coreConfigs))
	for _, coreConfig := range l.config.coreConfigs {
		cores = append(cores, zapcore.NewCore(coreConfig.Enc, coreConfig.Ws, coreConfig.Lvl))
	}

	logger := zap.New(
		zapcore.NewTee(cores[:]...),
		l.config.zapOpts...)

	l.l = logger
}

// Logger is used to return an instance of *klogzap.Logger for custom fields, etc.
func (l *Logger) Logger() *zap.Logger {
	return l.l
}

func (l *Logger) Sync() {
	_ = l.l.Sync()
}
