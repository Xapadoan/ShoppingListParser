package logger

import dom "github.com/Xapadoan/shplsprsr/logger/domain"

type ILogger = dom.ILogger

type ILoggerGateway interface {
	NewLogger(prefix string) ILogger
}
