package ingredient

import dom "github.com/Xapadoan/shplsprsr/ingredient/domain"

type IngredientUnit = dom.IngredientUnit

const (
	Undefined  IngredientUnit = dom.Undefined
	Gram                      = dom.Gram
	Milliliter                = dom.Milliliter
	Cas                       = dom.Cas
	Cac                       = dom.Cac
	Unit                      = dom.Unit
)

type IngredientQuantity = dom.IngredientQuantity
type IParseIngredientQuantity = dom.IParseIngredientQuantity

type IngredientError = dom.IngredientError

type IIngredientGateway interface {
	ParseIngredientQuantity(str string, parsers []IParseIngredientQuantity) (*IngredientQuantity, *dom.IngredientError)
}
