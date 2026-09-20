package domain

import (
	"strconv"
)

type IngredientQuantity struct {
	Name   string
	Amount float32
	Unit   Unit
}

func (q IngredientQuantity) String() string {
	return strconv.FormatFloat(float64(q.Amount), 'f', 2, 32) + q.Unit.String() + " " + q.Name
}
