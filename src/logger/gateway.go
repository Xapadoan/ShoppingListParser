package logger

import (
	app "github.com/Xapadoan/shplsprsr/logger/application"
	dom "github.com/Xapadoan/shplsprsr/logger/domain"
	infra "github.com/Xapadoan/shplsprsr/logger/infrastructure"
)

type LoggerGateway struct{}

func (g *LoggerGateway) NewLogger(prefix string) dom.ILogger {
	level := dom.Debug
	consoleLog := infra.NewConsoleLog(level)
	return app.NewLogger(
		level,
		consoleLog.Log,
		app.CompactFormatter,
		prefix,
	)
}
