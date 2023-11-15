package klogzap

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

// inArray check if a string in a slice
func inArray(key ExtraKey, arr []ExtraKey) bool {
	for _, k := range arr {
		if k == key {
			return true
		}
	}
	return false
}
func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}
func OtelSeverityText(lv zapcore.Level) string {
	s := lv.CapitalString()
	if s == "DPANIC" || s == "PANIC" {
		s = "FATAL"
	}
	return s
}
