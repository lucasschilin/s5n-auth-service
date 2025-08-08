package logger

import "fmt"

var WarningConfig LevelConfig = LevelConfig{
	weight: 60,
	label:  "WARNING",
}

func (l *logger) Warning(msg string) {
	if l.minLevel > WarningConfig.weight {
		return
	}

	createLog(2, WarningConfig.label, msg, "", l.tracertID).Pretty().Stdout()
}

func (l *logger) Warningf(formatMsg string, a ...any) {
	l.Warning(fmt.Sprintf(formatMsg, a...))
}
