package presentation

import (
	"encoding/json"

	dom "github.com/Xapadoan/shplsprsr/recipe/domain"
)

func RecipeAsJsonBody(r *dom.Recipe) ([]byte, error) {
	var ingredientsAsString []string
	for _, i := range r.Ingredients() {
		ingredientAsJson, jsonIngredientErr := json.Marshal(i)
		if jsonIngredientErr != nil {
			return nil, jsonIngredientErr
		}
		ingredientsAsString = append(ingredientsAsString, string(ingredientAsJson))
	}

	jsonRecipeBody, jsonRecipeError := json.Marshal(struct {
		Name        string
		Quantity    int
		Ingredients []string
		Steps       []string
	}{
		r.Name(),
		int(r.NumberOfPeopleEating()),
		ingredientsAsString,
		r.Steps(),
	})

	if jsonRecipeError != nil {
		return nil, jsonRecipeError
	}

	return jsonRecipeBody, nil
}
