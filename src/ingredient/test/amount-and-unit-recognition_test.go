package test

import (
	"math"
	"testing"

	dom "github.com/Xapadoan/shplsprsr/ingredient/domain"
)

func TestAmountAndUnitParser(t *testing.T) {
	for _, sample := range RealLifeSamples() {
		ingredientQuantity, parseErr := dom.ParseAmountAndUnit(sample.TestString)

		if parseErr != nil {
			t.Errorf("Failed to parse sample \"%v\":\n\t"+parseErr.Error(), sample.TestString)
			break
		}

		if math.Round(float64(ingredientQuantity.Amount)) != math.Round(float64(sample.ExpectedResult.Amount)) {
			t.Errorf(
				"Parsed amount does not match the expected one for sample \"%v\":\n\tExpected: %f, Received: %f",
				sample.TestString, sample.ExpectedResult.Amount, sample.ExpectedResult.Amount,
			)
		}

		if ingredientQuantity.Unit != sample.ExpectedResult.Unit {
			t.Errorf(
				"Parsed unit does not match the expected one for sample \"%v\":\n\tExpected: %v, Received: %v",
				sample.TestString, sample.ExpectedResult.Unit.String(), ingredientQuantity.Unit.String(),
			)
		}
	}
}
