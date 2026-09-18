package test

import (
	"fmt"
	"regexp"
	"testing"

	dom "github.com/Xapadoan/shplsprsr/ingredient/domain"
)

func TestUnitRegexes(t *testing.T) {
	for _, sample := range UnitRecognitionSamples() {
		fmt.Println(dom.UnitRecognitionRegex())
		regex := regexp.MustCompile(".*[[]Amount[]] ?(" + dom.UnitRecognitionRegex() + ").*")
		matches := regex.FindStringSubmatch(sample.TestString)
		if len(matches) < 1 && sample.ExpectedResult.Unit != dom.Unit_Unit {
			t.Errorf("Failed to detect unit for sample %v", sample.TestString)
		}
	}
}
