package domain

import (
	"slices"
	"strings"

	ing "github.com/Xapadoan/shplsprsr/ingredient"
	rcp "github.com/Xapadoan/shplsprsr/recipe/domain"
)

type ShoppingList map[string][]*ing.IngredientQuantity

func (shpls ShoppingList) Add(item *ing.IngredientQuantity) {
	value, exists := shpls[item.Name]
	if !exists {
		shpls[item.Name] = []*ing.IngredientQuantity{item}
		return
	}

	matchingUnitIndex := slices.IndexFunc(value, func(q *ing.IngredientQuantity) bool { return q.Unit == item.Unit })
	if matchingUnitIndex < 0 {
		value = append(value, item)
		return
	}

	value[matchingUnitIndex] = &ing.IngredientQuantity{
		Name:   value[matchingUnitIndex].Name,
		Amount: value[matchingUnitIndex].Amount + item.Amount,
		Unit:   value[matchingUnitIndex].Unit,
	}
}

func (shpl *ShoppingList) AddRecipeIngredients(
	recipe *rcp.Recipe,
) *ShoppingListError {
	for _, ing := range recipe.Ingredients() {
		shpl.Add(ing)
	}

	return nil
}

func (shpls ShoppingList) String() string {
	itemsStr := []string{}
	for key, value := range shpls {
		quantitiesStr := []string{}
		for _, v := range value {
			quantitiesStr = append(quantitiesStr, v.String())
		}
		itemsStr = append(itemsStr, key+":["+strings.Join(quantitiesStr, ",")+"]")
	}

	return "{" + strings.Join(itemsStr, ",") + "}"
}
