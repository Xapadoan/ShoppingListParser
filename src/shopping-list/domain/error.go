package domain

type ShoppingListErrorCode int

const (
	Unknown ShoppingListErrorCode = iota
	ParsingFailed
)

type ShoppingListError struct {
	Code    ShoppingListErrorCode
	Message string
}

func (e ShoppingListError) String() string {
	switch e.Code {
	case ParsingFailed:
		return "ParsingFailed"
	default:
		return "Unknown"
	}
}

func (e *ShoppingListError) Error() string {
	return e.String()
}
