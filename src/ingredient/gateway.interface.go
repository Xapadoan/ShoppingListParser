package ingredient

import dom "github.com/Xapadoan/shplsprsr/ingredient/domain"

type Unit = dom.Unit

const (
	Unit_Undefined  Unit = dom.Unit_Undefined
	Unit_Gram            = dom.Unit_Gram
	Unit_Milliliter      = dom.Unit_Milliliter
	Unit_Cas             = dom.Unit_Cas
	Unit_Cac             = dom.Unit_Cac
	Unit_Unit            = dom.Unit_Unit
)

type IngredientQuantity = dom.IngredientQuantity
type IParseIngredientQuantity = dom.IParseIngredientQuantity

type IngredientError = dom.IngredientError

type IIngredientGateway interface {
	ParseIngredientQuantity(str string) (*IngredientQuantity, *dom.IngredientError)
}
