package infrastructure

import (
	"fmt"
	"strings"

	dom "github.com/Xapadoan/shplsprsr/logger/domain"
)

type ConsoleLog struct {
	level dom.LogLevel
}

func NewConsoleLog(level dom.LogLevel) ConsoleLog {
	return ConsoleLog{level: level}
}

const reset = "\033[0m"
const red = "\033[31m"
const yellow = "\033[33m"
const white = "\033[97m"
const blue = "\033[34m"

func (logger *ConsoleLog) Log(level dom.LogLevel, msg ...string) {
	var color string
	switch level {
	case dom.Error:
		color = red
	case dom.Warn:
		color = yellow
	case dom.Info:
		color = white
	case dom.Debug:
		color = blue
	}

	fmt.Println(color + strings.Join(msg, " ") + reset)
}
