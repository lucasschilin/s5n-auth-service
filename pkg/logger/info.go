package logger

var InfoConfig LevelConfig = LevelConfig{
	weight: 40,
	label:  "INFO",
}

func (l *logger) Info(msg string) {
	if l.minLevel > InfoConfig.weight {
		return
	}

	createLog(2, InfoConfig.label, msg, "", l.tracertID).Pretty().Stdout()
}
