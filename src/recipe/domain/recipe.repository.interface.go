package domain

type IGetRecipe = func(id string) (*Recipe, *RecipeError)

type FindRecipesParams struct {
	Ids []string
}
type IFindRecipe interface {
	Find(params *FindRecipesParams) []*Recipe
}
