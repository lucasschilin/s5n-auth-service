package logger

import "fmt"

func (l *logString) Stdout() {
	fmt.Println(l.text)
}
