package domain

import (
	"regexp"
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

func ParseAmountAndUnit(text string) (*IngredientQuantity, *IngredientError) {
	regex := regexp.MustCompile("(" + AmountRecognitionRegexp() + ") ?(" + UnitRecognitionRegexp() + ")?")
	matches := regex.FindStringSubmatch(text)
	if len(matches) != 3 {
		return &IngredientQuantity{Amount: 1, Unit: Unit_Unit, Name: ""}, nil
	}

	amount, amountErr := ParseAmount(matches[1])
	if amountErr != nil {
		return nil, amountErr
	}

	unitAndRatio, unitAndRatioErr := ParseUnitAndRatio(matches[2])
	if unitAndRatioErr != nil {
		return nil, unitAndRatioErr
	}

	return &IngredientQuantity{Amount: amount * unitAndRatio.Ratio, Unit: unitAndRatio.Unit, Name: ""}, nil
}
