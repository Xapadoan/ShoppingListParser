package test

import (
	"regexp"
	"testing"

	dom "github.com/Xapadoan/shplsprsr/ingredient/domain"
)

func TestUnitRecognitionRegexp(t *testing.T) {
	for _, sample := range UnitRecognitionSamples() {
		regex := regexp.MustCompile(".*[[]Amount[]] ?(" + dom.UnitRecognitionRegexp() + ").*")
		matches := regex.FindStringSubmatch(sample.TestString)
		if len(matches) < 1 && sample.ExpectedResult.Unit != dom.Unit_Unit {
			t.Errorf("Failed to detect unit for sample \"%v\"", sample.TestString)
		}
	}
}
