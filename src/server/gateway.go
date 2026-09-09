package server

import (
	dom "github.com/Xapadoan/shplsprsr/server/domain"
	infra "github.com/Xapadoan/shplsprsr/server/infrastructure"

	log "github.com/Xapadoan/shplsprsr/logger"
)

type ServerGateway struct {
	loggerGateway log.ILoggerGateway
}

func NewServerGateway(loggerGateway log.ILoggerGateway) IServerGateway {
	return &ServerGateway{loggerGateway: loggerGateway}
}

func (g *ServerGateway) NewHttpServer() dom.IServeHttp {
	return infra.NewHttpMuxServer(g.loggerGateway.NewLogger("Server"))
}

const (
	BadRequest = dom.BadRequest
	NotFound   = dom.NotFound
	Internal   = dom.Internal
)

const (
	Ok              = dom.Ok
	BadRequestError = dom.BadRequestError
	NotFoundError   = dom.NotFoundError
	InternalError   = dom.InternalError
)

const (
	GET  = dom.GET
	POST = dom.POST
)
