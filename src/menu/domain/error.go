package domain

type MenuErrorCode int

const (
	Unknown MenuErrorCode = iota
	NotFound
	FetchFailed
	InvalidData
)

type MenuError struct {
	Code MenuErrorCode
}

func (e MenuError) String() string {
	switch e.Code {
	case NotFound:
		return "Not Found"
	case FetchFailed:
		return "Fetch Failed"
	case InvalidData:
		return "Invalid Data"
	default:
		return "Unknown"
	}
}
func (e *MenuError) Error() string {
	return e.String()
}
