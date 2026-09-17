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
	return strconv.Itoa(int(q.Amount)) + q.Unit.String() + " " + q.Name
}
