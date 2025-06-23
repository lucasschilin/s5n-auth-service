package logger

type LevelConfig struct {
	weight int
	label  string
}

func (l *logger) SetTracertID(tracertID string) {
	l.tracertID = tracertID
}
