package domain

type IParseIngredientQuantity = func(text string) (*IngredientQuantity, *IngredientError)
