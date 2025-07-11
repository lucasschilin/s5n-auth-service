package logger

import "fmt"

var DebugConfig LevelConfig = LevelConfig{
	weight: 20,
	label:  "DEBUG",
}

func (l *logger) Debug(msg string) {
	if l.minLevel > DebugConfig.weight {
		return
	}

	createLog(2, DebugConfig.label, msg, "", l.tracertID).Pretty().Stdout()
}

func (l *logger) Debugf(formatMsg string, a ...any) {
	l.Debug(fmt.Sprintf(formatMsg, a...))
}
