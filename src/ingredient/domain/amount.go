package domain

import (
	"regexp"
	"strconv"
	"strings"
)

func parseIntegerAmount(pattern string) (float32, error) {
	integer, err := strconv.Atoi(pattern)
	if err != nil {
		return 0, &IngredientError{ParsingFailed, "Atoi failed for pattern \"" + pattern + "\":\n" + err.Error()}
	}
	return float32(integer), nil
}

func parseDecimalAmount(pattern string) (float32, error) {
	integers := strings.Split(pattern, "/")
	if len(integers) != 2 {
		return 0, &IngredientError{ParsingFailed, "Pattern " + pattern + " does not match [int]/[int]"}
	}

	numerator, numeratorError := strconv.Atoi(strings.Trim(integers[0], " "))
	if numeratorError != nil {
		return 0, &IngredientError{ParsingFailed, "Atoi failed for pattern \"" + pattern + "\": \n" + numeratorError.Error()}
	}
	denumerator, denumeratorError := strconv.Atoi(strings.Trim(integers[1], " "))
	if numeratorError != nil {
		return 0, &IngredientError{ParsingFailed, "Atoi failed for pattern \"" + pattern + "\": \n" + denumeratorError.Error()}
	}
	return float32(numerator) / float32(denumerator), nil
}

func parseIntegerDecimalHybridAmount(pattern string) (float32, error) {
	splitRegexp := regexp.MustCompile("([0-9]+) (1 ?/ ?[0-9]{1})")
	matches := splitRegexp.FindStringSubmatch(pattern)

	if len(matches) < 3 {
		return 0, &IngredientError{ParsingFailed, "pattern \"" + pattern + "\" does not match [int] 1 / [int]"}
	}

	multiplier, multiplierErr := parseIntegerAmount(matches[1])
	if multiplierErr != nil {
		return 0, &IngredientError{ParsingFailed, multiplierErr.Error()}
	}

	decimal, decimalErr := parseDecimalAmount(matches[2])
	if decimalErr != nil {
		return 0, &IngredientError{ParsingFailed, decimalErr.Error()}
	}

	return multiplier * decimal, nil
}

func parseSpecificCharacterAmount(pattern string) (float32, error) {
	switch pattern {
	case "¼":
		return 0.25, nil
	case "½":
		return 0.5, nil
	default:
		return 0, &IngredientError{ParsingFailed, "pattern \"" + pattern + "\" is not among recognized specific characters"}
	}
}

func amountRecognitionMap() map[string]func(string) (float32, error) {
	return map[string]func(string) (float32, error){
		"[0-9]+ 1 ?/ ?[0-9]{1}": parseIntegerDecimalHybridAmount,
		"1 ?/ ?[0-9]{1}":        parseDecimalAmount,
		"[0-9]+":                parseIntegerAmount,
		"½|¼":                   parseSpecificCharacterAmount,
	}
}

func AmountRecognitionRegexp() string {
	var patterns []string

	for pattern := range amountRecognitionMap() {
		patterns = append(patterns, pattern)
	}

	return strings.Join(patterns, "|")
}
