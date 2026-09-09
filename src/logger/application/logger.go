package application

import (
	dom "github.com/Xapadoan/shplsprsr/logger/domain"
)

type Logger struct {
	level  dom.LogLevel
	log    dom.ILog
	format dom.IFormat
	prefix string
}

func NewLogger(level dom.LogLevel, log dom.ILog, format dom.IFormat, prefix string) *Logger {
	return &Logger{log: log, format: format, prefix: prefix}
}

func (l *Logger) print(level dom.LogLevel, msg ...string) {
	if level < l.level {
		return
	}

	l.log(level, l.format(level, msg...))
}

func (l *Logger) Error(err error, msg ...string) {
	l.print(dom.Error, msg...)
	l.print(dom.Error, err.Error())
}

func (l *Logger) Warn(msg ...string) {
	l.print(dom.Warn, msg...)
}

func (l *Logger) Info(msg ...string) {
	l.print(dom.Info, msg...)
}

func (l *Logger) Debug(msg ...string) {
	l.print(dom.Debug, msg...)
}
