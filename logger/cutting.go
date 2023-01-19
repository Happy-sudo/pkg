package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
	"strings"
)

type CuttingLogConfig struct {
	Filename   string `json:"file_name" json:"fileName"`     //路径
	MaxSize    int    `json:"max_size" json:"maxSize"`       //日志的最大大小（M）
	MaxBackups int    `json:"max_backups" json:"maxBackups"` //日志的最大保存数量
	MaxAge     int    `json:"max_age" json:"maxAge"`         //日志文件存储最大天数
	Compress   bool   `json:"compress"`                      //是否执行压缩
	LocalTime  bool   `json:"local_time" json:"localTime"`   //是否使用格式化时间辍
}

//CuttingLogWriter 切割日志
func (conf *CuttingLogConfig) CuttingLogWriter() (zapcore.WriteSyncer, error) {

	if conf.Filename == "" || strings.HasSuffix(conf.Filename, ".log") {
		return nil, errors.Wrap(errors.New("Cutting Log FileName error"), "Cutting Log FileName error")
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s", conf.Filename),
		MaxSize:    conf.MaxSize,    //日志的最大大小（M）
		MaxBackups: conf.MaxBackups, //日志的最大保存数量
		MaxAge:     conf.MaxAge,     //日志文件存储最大天数
		Compress:   conf.Compress,   //是否执行压缩
		LocalTime:  conf.LocalTime,  // 是否使用格式化时间辍
	}
	return zapcore.AddSync(lumberJackLogger), nil
}
