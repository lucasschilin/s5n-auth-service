package logger

import (
	"encoding/json"
	"fmt"
)

func (l *logStructure) Pretty() *logString {

	timestamptz := "--"
	if l.TimestampTZ != "" {
		timestamptz = l.TimestampTZ
	}

	level := "--"
	if l.Level != "" {
		level = l.Level
	}

	message := "--"
	if l.Message != "" {
		message = l.Message
	}

	error := "--"
	if l.Error != "" {
		error = "error:" + l.Error
	}

	fileLine := "--"
	if l.FileLine != "" {
		fileLine = l.FileLine
	}

	tracertID := "--"
	if l.TracertID != "" {
		tracertID = l.TracertID
	}

	text := fmt.Sprintf(
		"%s | %s | %s | %s | %s | %s",
		timestamptz, level, message, error, fileLine, tracertID,
	)
	return &logString{
		text: text,
	}
}
func (l *logStructure) JSON() *logString {
	encoded, err := json.Marshal(l)
	if err != nil {
		return &logString{
			text: "{}",
		}
	}

	return &logString{
		text: string(encoded),
	}
}
