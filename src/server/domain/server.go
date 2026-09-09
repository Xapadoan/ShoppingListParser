package domain

type IServeHttp interface {
	RegisterRoute(routes Route)
	Start(port int)
}
