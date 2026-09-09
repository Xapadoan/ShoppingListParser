package domain

type IngredientErrorCode uint8

const (
	Unknown IngredientErrorCode = iota
	ParsingFailed
)

type IngredientError struct {
	Code    IngredientErrorCode
	Message string
}

func (e IngredientError) String() string {
	switch e.Code {
	case ParsingFailed:
		return "Parsing Failed"
	default:
		return "Unknown"
	}
}

func (e *IngredientError) Error() string {
	return e.String()
}
