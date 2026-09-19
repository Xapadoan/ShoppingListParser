package domain

import (
	"encoding/json"

	ing "github.com/Xapadoan/shplsprsr/ingredient/domain"
)

type Recipe struct {
	name                 string
	numberOfPeopleEating uint8
	ingredients          []*ing.IngredientQuantity
	steps                []string
}

func NewRecipe(
	name string,
	numberOfPeopleEating uint8,
	ingredients []*ing.IngredientQuantity,
	steps []string,
) *Recipe {
	return &Recipe{name, numberOfPeopleEating, ingredients, steps}
}

func (r *Recipe) Name() string {
	return r.name
}

func (r *Recipe) NumberOfPeopleEating() uint8 {
	return r.numberOfPeopleEating
}

func (r *Recipe) Ingredients() []*ing.IngredientQuantity {
	return r.ingredients
}

func (r *Recipe) Steps() []string {
	return r.steps
}

func (r *Recipe) AdaptQuantity(numberOfPeopleEating uint16) *Recipe {
	recipe := Recipe{r.name, r.numberOfPeopleEating, r.ingredients, r.steps}
	for _, ing := range recipe.ingredients {
		ing.Amount = ing.Amount * float32(numberOfPeopleEating) / float32(r.numberOfPeopleEating)
	}

	return &recipe
}

func (r *Recipe) MarshalJSON() ([]byte, error) {

	var ingredientsAsString []string
	for _, i := range r.Ingredients() {
		ingredientsAsString = append(ingredientsAsString, i.String())
	}
	return json.Marshal(struct {
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
}
