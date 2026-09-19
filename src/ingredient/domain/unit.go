package domain

import "strings"

type Unit int

const (
	Unit_Undefined Unit = iota
	Unit_Gram
	Unit_Unit
	Unit_Cas
	Unit_Cac
	Unit_Milliliter
)

func (u Unit) String() string {
	switch u {
	case Unit_Undefined:
		return "?"
	case Unit_Gram:
		return "g"
	case Unit_Unit:
		return ""
	case Unit_Cas:
		return "cas"
	case Unit_Cac:
		return "cac"
	case Unit_Milliliter:
		return "ml"
	}

	return "undefined"
}

type UnitConversionRatio struct {
	Unit  Unit
	Ratio float32
}

type UnitAndRatioRecognition struct {
	pattern      string
	unitAndRatio UnitConversionRatio
}

func unitAndRatioRecognitionArray() []UnitAndRatioRecognition {
	return []UnitAndRatioRecognition{
		// Unit_Cac
		{"cac", UnitConversionRatio{Unit_Cac, 1}},
		{"cuillère à café", UnitConversionRatio{Unit_Cac, 1}},
		{"c. à café", UnitConversionRatio{Unit_Cac, 1}},

		// Unit_Cas
		{"cas", UnitConversionRatio{Unit_Cas, 1}},
		{"cuillère à soupe", UnitConversionRatio{Unit_Cas, 1}},
		{"cuillères à soupe", UnitConversionRatio{Unit_Cas, 1}},
		{"c. à soupe", UnitConversionRatio{Unit_Cas, 1}},

		// Unit_Gram
		{"g", UnitConversionRatio{Unit_Gram, 1}},

		// Unit_Milliliter
		{"ml", UnitConversionRatio{Unit_Milliliter, 1}},

		// Unit_Unit
		{"u", UnitConversionRatio{Unit_Unit, 1}},
		{"", UnitConversionRatio{Unit_Unit, 1}},
	}
}

func UnitRecognitionRegexp() string {
	var patterns []string

	for _, unitAndRatioRecognition := range unitAndRatioRecognitionArray() {
		patterns = append(patterns, unitAndRatioRecognition.pattern)
	}

	return strings.Join(patterns, "|")
}

func (u *Unit) MarshalJSON() ([]byte, error) {
	return []byte("\"" + u.String() + "\""), nil
}
