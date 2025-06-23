package logger

import (
	"fmt"
	"runtime"
	"time"
)

type logger struct {
	minLevel  int
	tracertID string
}

type logStructure struct {
	Level       string `json:"level"`
	TimestampTZ string `json:"timestamptz"`
	Message     string `json:"message"`
	FileLine    string `json:"fileline"`
	Error       string `json:"error,omitempty"`
	TracertID   string `json:"tracert_id,omitempty"`
}

type logString struct {
	text string
}

type Logger interface {
	Error(error error, msg string)
	Errorf(error error, formatMsg string, a ...any)
	Warning(msg string)
	Warningf(formatMsg string, a ...any)
	Info(msg string)
	Infof(formatMsg string, a ...any)
	Debug(msg string)
	Debugf(formatMsg string, a ...any)
}

func New(level int) Logger {
	return &logger{
		minLevel: level,
	}
}

func createLog(
	depth int, level, message, errorMessage, tracertID string,
) *logStructure {

	_, file, line, ok := runtime.Caller(depth)
	var fileLine string
	if ok {
		fileLine = fmt.Sprintf("%s:%d", file, line)
	}

	return &logStructure{
		TimestampTZ: time.Now().Format(time.RFC3339),
		Level:       level,
		Message:     message,
		Error:       errorMessage,
		TracertID:   tracertID,
		FileLine:    fileLine,
	}
}
