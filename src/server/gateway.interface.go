package server

import dom "github.com/Xapadoan/shplsprsr/server/domain"

type IServerGateway interface {
	NewHttpServer() dom.IServeHttp
}

type IServeHttp = dom.IServeHttp
type Route = dom.Route
type RouteRequest = dom.RouteRequest
type RouteResponse = dom.RouteResponse
type ServerError = dom.ServerError
