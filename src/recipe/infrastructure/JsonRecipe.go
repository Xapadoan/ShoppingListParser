package infrastructure

import (
	"regexp"
	"strconv"
	"strings"

	dom "github.com/Xapadoan/shplsprsr/recipe/domain"

	ing "github.com/Xapadoan/shplsprsr/ingredient"
)

type JsonRecipe struct {
	Name        string
	Quantity    int
	Ingredients []string
	Steps       []string
}

func (r *JsonRecipe) DomainAdapter() (*dom.Recipe, *dom.RecipeError) {
	if r.Quantity > 255 {
		return &dom.Recipe{}, dom.NewRecipeError(dom.InvalidData, "Quantity is > 255")
	}

	var ingredients []*ing.IngredientQuantity
	for _, ingStr := range r.Ingredients {
		ing, ingErr := ingredientParser(ingStr)
		if ingErr != nil {
			return &dom.Recipe{}, dom.NewRecipeError(dom.InvalidData, "Failed to parse ingredient "+ingStr+":\n"+ingErr.Error())
		}
		ingredients = append(ingredients, ing)
	}

	return dom.NewRecipe(r.Name, uint8(r.Quantity), ingredients, r.Steps), nil
}

type JsonIngredientParser = func(text string) (*ing.IngredientQuantity, *dom.RecipeError)

func ingredientParser(text string) (*ing.IngredientQuantity, *dom.RecipeError) {
	for _, parse := range []JsonIngredientParser{amountUnitNameParser, nameAmountUnitName, nameOnlyParser} {
		i, err := parse(text)
		if err == nil {
			return i, nil
		}
	}

	return nil, dom.NewRecipeError(dom.InvalidData, "Failed to parse ingredient"+text)
}

func nameOnlyParser(text string) (*ing.IngredientQuantity, *dom.RecipeError) {
	return &ing.IngredientQuantity{Name: strings.ToLower(text), Amount: 1}, nil
}

func nameAmountUnitName(text string) (*ing.IngredientQuantity, *dom.RecipeError) {
	re := regexp.MustCompile("^(.+)[- ]+([0-9]+) ?(kg|g|u|cas|cac|ml|mL|cl|cL)?$")
	matches := re.FindStringSubmatch(text)
	if len(matches) < 1 {
		msg := "String" + text + "does not match regexp" + re.String()
		return nil, dom.NewRecipeError(dom.InvalidData, msg)
	}

	quantity, quantityErr := strconv.ParseUint(matches[2], 10, 64)
	if quantityErr != nil {
		msg := "Failed to parse" + matches[2] + "as integer:\n" + quantityErr.Error()
		return nil, dom.NewRecipeError(dom.InvalidData, msg)
	}

	unit, ratio := parseUnitAndRatio(matches[3])

	return &ing.IngredientQuantity{Name: strings.ToLower(matches[1]), Amount: uint16(quantity) * ratio, Unit: unit}, nil
}

func amountUnitNameParser(text string) (*ing.IngredientQuantity, *dom.RecipeError) {
	re := regexp.MustCompile("^([0-9]+) ?(kg|g|u|cas|cac|ml|mL|cl|cL)? (de )?(.*)$")
	matches := re.FindStringSubmatch(strings.ToLower(text))
	if len(matches) < 1 {
		msg := "String" + text + "does not match regexp" + re.String()
		return nil, dom.NewRecipeError(dom.InvalidData, msg)
	}

	quantity, quantityErr := strconv.ParseUint(matches[1], 10, 64)
	if quantityErr != nil {
		msg := "Failed to parse" + matches[1] + "as integer:\n" + quantityErr.Error()
		return nil, dom.NewRecipeError(dom.InvalidData, msg)
	}

	unit, ratio := parseUnitAndRatio(matches[2])

	return &ing.IngredientQuantity{Name: strings.ToLower(matches[4]), Amount: uint16(quantity) * ratio, Unit: unit}, nil
}

func parseUnitAndRatio(text string) (ing.IngredientUnit, uint16) {
	switch text {
	case "kg":
		return ing.Gram, 1000
	case "g":
		return ing.Gram, 1
	case "ml", "mL":
		return ing.Milliliter, 1
	case "cl", "cL":
		return ing.Milliliter, 10
	case "cas", "Cas":
		return ing.Cas, 1
	case "cac", "Cac":
		return ing.Cac, 1
	default:
		return ing.Unit, 1
	}
}
