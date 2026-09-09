package domain

type ServerErrorCode int

const (
	BadRequest ServerErrorCode = iota
	NotFound
	Internal
)

type ServerError struct {
	Code ServerErrorCode
}

func (e ServerError) String() string {
	switch e.Code {
	case BadRequest:
		return "Bad Request"
	case NotFound:
		return "Not Found"
	case Internal:
		return "Internal"
	default:
		return "Internal"
	}
}
