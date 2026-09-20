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

func (c IngredientErrorCode) String() string {
	switch c {
	case ParsingFailed:
		return "Parsing Failed"
	default:
		return "Unknown"
	}
}

func (e IngredientError) String() string {
	return "[" + e.Code.String() + "] " + e.Message
}

func (e *IngredientError) Error() string {
	return e.String()
}
