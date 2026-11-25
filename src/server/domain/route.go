package domain

type StatusCode int

const (
	Ok              StatusCode = 200
	Created         StatusCode = 201
	NoContent       StatusCode = 204
	InternalError   StatusCode = 500
	NotFoundError   StatusCode = 404
	BadRequestError StatusCode = 400
)

func StatusCodeFromServerErrorCode(code ServerErrorCode) StatusCode {
	switch code {
	case Internal:
		return InternalError
	case NotFound:
		return NotFoundError
	case BadRequest:
		return BadRequestError
	default:
		return InternalError
	}
}

type Method string

const (
	GET  Method = "GET"
	POST Method = "POST"
	PUT  Method = "PUT"
)

func (m Method) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	default:
		return "GET"
	}
}

type RouteResponse struct {
	Status StatusCode
	Body   []byte
}

type IHandleRequest func(request *RouteRequest, response *RouteResponse) *ServerError

type Route struct {
	Method        Method
	Path          string
	HandleRequest IHandleRequest
	ParseURL      IParseURL
}
type IParseURL func(url string, req *RouteRequest) *ServerError
