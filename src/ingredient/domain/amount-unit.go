package domain

func parseAmountAndUnitRegexp() string {
	return "(" + AmountRecognitionRegexp() + ")(?: ?(" + UnitRecognitionRegexp() + "))?"
}

func parseAmountAndUnitRecognizedPatterns(recognizedAmountPattern string, recognizedUnitPattern string) (*IngredientQuantity, *IngredientError) {
	amount, amountErr := ParseAmount(recognizedAmountPattern)
	if amountErr != nil {
		return nil, amountErr
	}

	unitAndRatio, unitAndRatioErr := ParseUnitAndRatio(recognizedUnitPattern)
	if unitAndRatioErr != nil {
		return nil, unitAndRatioErr
	}

	return &IngredientQuantity{Amount: amount * unitAndRatio.Ratio, Unit: unitAndRatio.Unit, Name: ""}, nil
}
