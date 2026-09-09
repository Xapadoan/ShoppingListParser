package domain

type RecipeErrorCode int

const (
	Unknown RecipeErrorCode = iota
	FetchFailed
	InvalidData
	NotFound
)

type RecipeError struct {
	code    RecipeErrorCode
	message string
}

func NewRecipeError(code RecipeErrorCode, message string) *RecipeError {
	return &RecipeError{code, message}
}

func (e *RecipeError) IsNotFound() bool {
	return e.code == NotFound
}

func (e RecipeError) String() string {
	switch e.code {
	case FetchFailed:
		return "Fetch Failed: " + e.message
	case InvalidData:
		return "Invalid Data: " + e.message
	case NotFound:
		return "Not Found: " + e.message
	default:
		return "Unknown: " + e.message
	}
}

func (e *RecipeError) Error() string {
	return e.String()
}
