package logger

var ErrorConfig LevelConfig = LevelConfig{
	weight: 90,
	label:  "ERROR",
}

func (l *logger) Error(error error, msg string) {
	if l.minLevel > ErrorConfig.weight {
		return
	}

	createLog(2, ErrorConfig.label, msg, error.Error(), l.tracertID).Pretty().Stdout()
}
