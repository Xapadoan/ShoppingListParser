package test

import (
	"math"
	"testing"

	dom "github.com/Xapadoan/shplsprsr/ingredient/domain"
)

func TestParseIngredientQuantity(t *testing.T) {
	for _, sample := range RealLifeSamples() {
		ingredient, parseErr := dom.ParseIngredientQuantity(sample.TestString)

		if parseErr != nil {
			t.Errorf("Parsing sample \"%v\" failed:\n%v\n", sample.TestString, parseErr.Error())
			continue
		}

		if ingredient.Unit != sample.ExpectedResult.Unit {
			t.Errorf(
				"Parsed unit is not the expected one for sample \"%v\":\nExpected: \"%v\", Received: \"%v\"",
				sample.TestString, sample.ExpectedResult.Unit.String(), ingredient.Unit.String(),
			)
			continue
		}
		if math.Round(float64(ingredient.Amount)) != math.Round(float64(sample.ExpectedResult.Amount)) {
			t.Errorf(
				"Parsed amount is not the expected one for sample \"%v\":\nExpected: \"%f\", Received: \"%f\"",
				sample.TestString, sample.ExpectedResult.Amount, ingredient.Amount,
			)
		}
	}
}
