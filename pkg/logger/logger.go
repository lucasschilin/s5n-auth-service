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

func New(level int) *logger {
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
