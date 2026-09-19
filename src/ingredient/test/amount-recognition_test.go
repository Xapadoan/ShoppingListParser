package test

import (
	"regexp"
	"testing"

	dom "github.com/Xapadoan/shplsprsr/ingredient/domain"
)

func TestAmountRecognitionRegexp(t *testing.T) {
	for _, sample := range RealLifeSamples() {
		regex := regexp.MustCompile("(" + dom.AmountRecognitionRegexp() + ")")
		matches := regex.FindStringSubmatch(sample.TestString)
		if len(matches) < 2 && sample.ExpectedResult.Amount != 1 {
			t.Errorf("Failed to detect any amount for sample \"%v\"", sample.TestString)
		}
	}
}

func TestAmountParsing(t *testing.T) {
	for _, sample := range AmountRecognizedPatternSamples() {
		parsedAmount, parsedAmountErr := dom.ParseAmount(sample.TestString)
		if parsedAmountErr != nil {
			t.Errorf("Failed to parse amount for sample \"%v\"", sample)
			continue
		}

		if parsedAmount != sample.ExpectedResult.Amount {
			t.Errorf(
				"Unexpected parsed amount for sample \"%v\"\n\tExpected %f, Received %f",
				sample.TestString,
				sample.ExpectedResult.Amount,
				parsedAmount,
			)
		}
	}
}
