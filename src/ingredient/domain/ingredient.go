package domain

import (
	"strconv"
)

type IngredientUnit int

const (
	Undefined IngredientUnit = iota
	Gram
	Unit
	Cas
	Cac
	Milliliter
)

func (u IngredientUnit) String() string {
	switch u {
	case Undefined:
		return "?"
	case Gram:
		return "g"
	case Unit:
		return ""
	case Cas:
		return "cas"
	case Cac:
		return "cac"
	case Milliliter:
		return "ml"
	}

	return "unknown"
}

func (u *IngredientUnit) MarshalJSON() ([]byte, error) {
	return []byte("\"" + u.String() + "\""), nil
}

type IngredientQuantity struct {
	Name   string
	Amount uint16
	Unit   IngredientUnit
}

func (q IngredientQuantity) String() string {
	return strconv.Itoa(int(q.Amount)) + q.Unit.String() + " " + q.Name
}
