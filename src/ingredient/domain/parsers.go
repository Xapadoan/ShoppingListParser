package domain

import (
	"regexp"
	"strings"
)

const CLUTTER_PATTERN string = "(?:(?:de|du|-|d') )?"

func parseNameOnly(text string) (*IngredientQuantity, *IngredientError) {
	return &IngredientQuantity{Name: strings.ToLower(text), Amount: 1, Unit: Unit_Unit}, nil
}

func parseNameAmountUnit(text string) (*IngredientQuantity, *IngredientError) {
	re := regexp.MustCompile("^(.+)[- ]+" + parseAmountAndUnitRegexp() + "$")
	matches := re.FindStringSubmatch(text)
	if len(matches) != 4 {
		msg := "String \n" + text + "\" does not match regexp \"" + re.String() + "\""
		return nil, &IngredientError{ParsingFailed, msg}
	}

	ingredientQuantity, parseErr := parseAmountAndUnitRecognizedPatterns(
		strings.Trim(matches[2], " "),
		strings.Trim(matches[3], " "),
	)
	if parseErr != nil {
		return nil, parseErr
	}

	ingredientQuantity.Name = strings.ToLower(matches[1])

	return ingredientQuantity, nil
}

func parseAmountUnitName(text string) (*IngredientQuantity, *IngredientError) {
	re := regexp.MustCompile("^" + parseAmountAndUnitRegexp() + " " + CLUTTER_PATTERN + "(.*)$")
	matches := re.FindStringSubmatch(strings.ToLower(text))
	if len(matches) != 4 {
		msg := "String \"" + text + "\" does not match regexp \"" + re.String() + "\""
		return nil, &IngredientError{ParsingFailed, msg}
	}

	ingredientQuantity, parseErr := parseAmountAndUnitRecognizedPatterns(
		strings.Trim(matches[1], " "),
		strings.Trim(matches[2], " "),
	)
	if parseErr != nil {
		return nil, parseErr
	}

	ingredientQuantity.Name = strings.ToLower(matches[3])

	return ingredientQuantity, nil
}

func ParseIngredientQuantity(text string) (*IngredientQuantity, *IngredientError) {
	parsers := []IParseIngredientQuantity{
		parseAmountUnitName,
		parseNameAmountUnit,
		parseNameOnly,
	}

	for _, parse := range parsers {
		i, err := parse(text)
		if err == nil {
			return i, nil
		}
	}

	return nil, &IngredientError{ParsingFailed, "All parsers failed"}
}
