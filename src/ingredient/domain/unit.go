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

type unitConversionRatio struct {
	unit  Unit
	ratio float32
}

func unitStringRecognitionMap() map[string]unitConversionRatio {
	return map[string]unitConversionRatio{
		// Unit_Cac
		"cac":             {Unit_Cac, 1},
		"cuillère à café": {Unit_Cac, 1},
		"c. à café":       {Unit_Cac, 1},

		// Unit_Cas
		"cas":               {Unit_Cas, 1},
		"cuillère à soupe":  {Unit_Cas, 1},
		"cuillères à soupe": {Unit_Cas, 1},

		// Unit_Gram
		"g": {Unit_Gram, 1},

		// Unit_Milliliter
		"ml": {Unit_Milliliter, 1},

		// Unit_Unit
		"u": {Unit_Unit, 1},
	}
}

func UnitRecognitionRegexp() string {
	var patterns []string

	for key := range unitStringRecognitionMap() {
		patterns = append(patterns, key)
	}

	return strings.Join(patterns, "|")
}

func (u *Unit) MarshalJSON() ([]byte, error) {
	return []byte("\"" + u.String() + "\""), nil
}
